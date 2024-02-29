// Info de los modulos
#include <linux/module.h>
// Info del kernel en tiempo real
#include <linux/kernel.h>
#include <linux/sched.h>

// Headers para modulos
#include <linux/init.h>
// Header necesario para proc_fs
#include <linux/proc_fs.h>
// Para dar acceso al usuario
#include <asm/uaccess.h>
// Para manejar el directorio /proc
#include <linux/seq_file.h>
// Para get_mm_rss
#include <linux/mm.h>

struct task_struct *cpu; // Estructura que almacena info del cpu

// Almacena los procesos
struct list_head *lstProcess;
// Estructura que almacena info de los procesos hijos
struct task_struct *child;
unsigned long rss;

MODULE_LICENSE("GPL");
MODULE_DESCRIPTION("Módulo CPU - Laboratorio Sistemas Operativos 1");
MODULE_AUTHOR("Carlos Javier Martinez Polanco");

static int mostrar_informacion_cpu(struct seq_file *archivo, void *v) {
    seq_printf(archivo, "{\n");

    seq_printf(archivo, "  \"procesos\": [\n");

    for_each_process(cpu) {
        seq_printf(archivo, "    {\n");
        seq_printf(archivo, "      \"PID\": %d,\n", cpu->pid);
        seq_printf(archivo, "      \"Nombre\": \"%s\",\n", cpu->comm);
        seq_printf(archivo, "      \"Estado\": %u,\n", cpu->__state);

        if (cpu->mm) {
            rss = get_mm_rss(cpu->mm) << PAGE_SHIFT;
            seq_printf(archivo, "      \"RSS\": %lu,\n", rss);
        } else {
            seq_printf(archivo, "      \"RSS\": null,\n");
        }

        seq_printf(archivo, "      \"UID\": %u\n", from_kuid(&init_user_ns, cpu->cred->user->uid));

        list_for_each(lstProcess, &(cpu->children)) {
            child = list_entry(lstProcess, struct task_struct, sibling);
            seq_printf(archivo, "      \"Child\": {\n");
            seq_printf(archivo, "        \"PID\": %d,\n", child->pid);
            seq_printf(archivo, "        \"Nombre\": \"%s\",\n", child->comm);
            seq_printf(archivo, "        \"Estado\": %u,\n", child->__state);

            if (child->mm) {
                rss = get_mm_rss(child->mm) << PAGE_SHIFT;
                seq_printf(archivo, "        \"RSS\": %lu,\n", rss);
            } else {
                seq_printf(archivo, "        \"RSS\": null,\n");
            }

            seq_printf(archivo, "        \"UID\": %u\n", from_kuid(&init_user_ns, child->cred->user->uid));
            seq_printf(archivo, "      }\n");
        }

        seq_printf(archivo, "    },\n");
    }

    seq_printf(archivo, "  ]\n");
    seq_printf(archivo, "}\n");

    return 0;
}

// Funcion que se ejecutara cada vez que se lea el archivo con el comando CAT
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

// Funcion a ejecuta al insertar el modulo en el kernel con insmod
static int __init cargar_modulo(void)
{
    proc_create("cpu_so1_1s2024", 0, NULL, &operaciones);
    printk(KERN_INFO "Módulo CPU cargado exitosamente\n");
    return 0;
}

// Funcion a ejecuta al remover el modulo del kernel con rmmod
static void descargar_modulo(void)
{
    remove_proc_entry("cpu_so1_1s2024", NULL);
    printk(KERN_INFO "Módulo RAM descargado exitosamente\n");
}

module_init(cargar_modulo);
module_exit(descargar_modulo);
