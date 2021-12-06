use actix_web::{get, post, HttpResponse, Responder};

#[get("/clubs")]
pub async fn users() -> impl Responder {
    HttpResponse::Ok().body("This is api/clubs")
}

#[post("/clubs")]
pub async fn create_user() -> impl Responder {
    HttpResponse::Ok().body("This is a POST to api/clubs")
}
