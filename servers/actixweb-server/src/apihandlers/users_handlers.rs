use actix_web::{get, post, HttpResponse, Responder};

#[get("/users")]
pub async fn users() -> impl Responder {
    HttpResponse::Ok().body("This is api/users")
}

#[post("/users")]
pub async fn create_user() -> impl Responder {
    HttpResponse::Ok().body("This is a POST to api/users")
}
