use actix_web::{web, App, HttpServer};

mod handlers;
mod apihandlers;

#[actix_web::main]
async fn main() -> std::io::Result<()> {
    HttpServer::new(|| {
        App::new()
            .service(handlers::index)
            .service(
                web::scope("/api")
                    .service(apihandlers::users)
                    .service(apihandlers::create_user)
            )
            .route("/hey", web::get().to(handlers::manual_hello))
            .default_service(
                web::route().to(handlers::not_found)
            )
    })
    .bind("127.0.0.1:8080")?
    .run()
    .await
}