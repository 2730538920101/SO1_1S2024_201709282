use hyper::service::{make_service_fn, service_fn};
use hyper::{Body, Request, Response, Server};

// Definir las constantes directamente en el archivo main.rs
const SERVER_HOST: &str = "127.0.0.1";
const SERVER_PORT: &str = "8080";

async fn receive_message(_req: Request<Body>) -> Result<Response<Body>, hyper::Error> {
    println!("Received GET request");

    let message = "Message received";
    let response_body = format!(r#"{{"message": "{}"}}"#, message);

    Ok(Response::new(Body::from(response_body)))
}

async fn insert_band(req: Request<Body>) -> Result<Response<Body>, hyper::Error> {
    let whole_body = hyper::body::to_bytes(req.into_body()).await?;
    println!("Received POST request: {:?}", whole_body);
    Ok(Response::new(Body::from(whole_body)))
}

#[tokio::main(flavor = "current_thread")]
async fn main() -> Result<(), Box<dyn std::error::Error>> {
    let host = SERVER_HOST;
    let port = SERVER_PORT;

    let addr = format!("{}:{}", host, port).parse()?;

    let make_svc = make_service_fn(|_conn| {
        async { Ok::<_, hyper::Error>(service_fn(router)) }
    });

    let server = Server::bind(&addr).serve(make_svc);

    println!("Listening on http://{}", addr);

    if let Err(e) = server.await {
        eprintln!("server error: {}", e);
    }

    Ok(())
}

async fn router(req: Request<Body>) -> Result<Response<Body>, hyper::Error> {
    match (req.method(), req.uri().path()) {
        (&hyper::Method::GET, "/receive") => receive_message(req).await,
        (&hyper::Method::POST, "/insert") => insert_band(req).await,
        _ => {
            let not_found = Response::builder()
                .status(hyper::StatusCode::NOT_FOUND)
                .body(Body::empty())
                .unwrap();
            Ok(not_found)
        }
    }
}
