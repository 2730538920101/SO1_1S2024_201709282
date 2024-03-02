// Info de los módulos
#include <linux/module.h>
#include <linux/kernel.h>
#include <linux/sched.h>
#include <linux/timekeeping.h>
#include <linux/timer.h>
#include <linux/jiffies.h>
#include <linux/proc_fs.h>
#include <linux/seq_file.h>
#include <asm/uaccess.h>
#include <linux/mm.h>
#include <linux/sched/cputime.h>
#include <linux/timekeeping.h>
#include <linux/time.h>

struct task_struct *cpu; // Estructura que almacena info del cpu
unsigned long rss;
// Declaración de la variable global para almacenar el valor de jiffies al inicio
unsigned long jiffies_at_boot;
MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Módulo CPU - Laboratorio Sistemas Operativos 1");
MODULE_AUTHOR("Carlos Javier Martinez Polanco");

static void imprimir_hijos(struct seq_file *archivo, struct list_head *hijos);
static unsigned long obtener_porcentaje_utilizacion_cpu(void);

static int mostrar_informacion_cpu(struct seq_file *archivo, void *v) {
    seq_printf(archivo, "{\n");
    unsigned long porcentaje_utilizacion = obtener_porcentaje_utilizacion_cpu();
    unsigned long decimal_part = porcentaje_utilizacion % 100;  // Obtener la parte decimal

    seq_printf(archivo, "  \"porcentaje_utilizacion_cpu\": %lu.%02lu,\n", porcentaje_utilizacion / 100, decimal_part);
    seq_printf(archivo, "  \"procesos\": [\n");

    int first_process = 1;

    for_each_process(cpu) {
        if (!first_process) {
            seq_printf(archivo, ",\n");
        }
        first_process = 0;

        seq_printf(archivo, "    {\n");
        seq_printf(archivo, "      \"ID_PROCESO\": %d,\n", cpu->pid);
        seq_printf(archivo, "      \"PID\": %d,\n", cpu->pid);
        seq_printf(archivo, "      \"Nombre\": \"%s\",\n", cpu->comm);
        seq_printf(archivo, "      \"Estado\": %u,\n", cpu->__state);

        if (cpu->mm) {
            rss = get_mm_rss(cpu->mm) << PAGE_SHIFT;
            seq_printf(archivo, "      \"RSS\": %lu,\n", rss);
        } else {
            seq_printf(archivo, "      \"RSS\": null,\n");
        }

        seq_printf(archivo, "      \"UID\": %u,\n", from_kuid(&init_user_ns, cpu->cred->user->uid));

        // Verificar si el proceso tiene hijos
        if (!list_empty(&cpu->children)) {
            seq_printf(archivo, "      \"hijos\": [\n");
            imprimir_hijos(archivo, &cpu->children);
            seq_printf(archivo, "\n      ]");
        } else {
            seq_printf(archivo, "      \"hijos\": []");
        }

        seq_printf(archivo, "\n    }");
    }

    seq_printf(archivo, "\n  ]\n");
    seq_printf(archivo, "}\n");

    return 0;
}

static void imprimir_hijos(struct seq_file *archivo, struct list_head *hijos) {
    struct list_head *lstProcess;
    struct task_struct *child;
    int first_child = 1;

    list_for_each(lstProcess, hijos) {
        if (!first_child) {
            seq_printf(archivo, ",\n");
        }
        first_child = 0;

        child = list_entry(lstProcess, struct task_struct, sibling);

        seq_printf(archivo, "        {\n");
        seq_printf(archivo, "          \"ID_PROCESO_HIJO\": %d,\n", child->pid);
        seq_printf(archivo, "          \"PID_HIJO\": %d,\n", child->pid);
        seq_printf(archivo, "          \"Nombre_HIJO\": \"%s\",\n", child->comm);
        seq_printf(archivo, "          \"Estado_HIJO\": %u,\n", child->__state);

        if (child->mm) {
            rss = get_mm_rss(child->mm) << PAGE_SHIFT;
            seq_printf(archivo, "          \"RSS_HIJO\": %lu,\n", rss);
        } else {
            seq_printf(archivo, "          \"RSS_HIJO\": null,\n");
        }

        seq_printf(archivo, "          \"UID_HIJO\": %u\n", from_kuid(&init_user_ns, child->cred->user->uid));
        seq_printf(archivo, "        }");
    }
}

// Funcion para obtener el porcentaje de utilizacion total del CPU
static unsigned long obtener_porcentaje_utilizacion_cpu(void) {
    unsigned long system_time = 0;
    unsigned long total_time = 0;

    for_each_process(cpu) {
        total_time += cpu->utime + cpu->stime;
        system_time += cpu->stime;
    }

    unsigned long elapsed_time = jiffies_to_usecs(jiffies - jiffies_at_boot);
    
    // Asegurarse de no dividir por cero
    if (elapsed_time == 0) {
        return 0;
    }

    // Calcular el porcentaje de utilización respecto al tiempo total
    unsigned long porcentaje_utilizacion = (system_time * 100) / elapsed_time;

    // Asegurarse de que el porcentaje esté en el rango [0, 100]
    if (porcentaje_utilizacion > 100) {
        porcentaje_utilizacion = 100;
    }

    printk(KERN_INFO "total_time: %lu\n", total_time);
    printk(KERN_INFO "system_time: %lu\n", system_time);
    printk(KERN_INFO "elapsed_time: %lu\n", elapsed_time);
    printk(KERN_INFO "porcentaje_utilizacion: %lu\n", porcentaje_utilizacion);

    return porcentaje_utilizacion;
}

// Función que se ejecutará cada vez que se lea el archivo con el comando CAT
static int abrir_archivo(struct inode *inode, struct file *file)
{
    return single_open(file, mostrar_informacion_cpu, NULL);
}

// Si el kernel es 5.6 o mayor se usa la estructura proc_ops
static struct proc_ops operaciones =
{
    .proc_open = abrir_archivo,
    .proc_read = seq_read
};

// Función a ejecutar al insertar el módulo en el kernel con insmod
static int __init cargar_modulo(void)
{
    proc_create("cpu_so1_1s2024", 0, NULL, &operaciones);
    jiffies_at_boot = jiffies;  // Inicializar jiffies_at_boot
    printk(KERN_INFO "Módulo CPU cargado exitosamente\n");
    return 0;
}

// Función a ejecutar al remover el módulo del kernel con rmmod
static void descargar_modulo(void)
{
    remove_proc_entry("cpu_so1_1s2024", NULL);
    printk(KERN_INFO "Módulo RAM descargado exitosamente\n");
}

module_init(cargar_modulo);
module_exit(descargar_modulo);
