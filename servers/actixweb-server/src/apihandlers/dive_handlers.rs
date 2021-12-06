use actix_web::{get, post, HttpResponse, Responder};

#[get("/dives")]
pub async fn users() -> impl Responder {
    HttpResponse::Ok().body("This is api/dives")
}

#[post("/dives")]
pub async fn create_user() -> impl Responder {
    HttpResponse::Ok().body("This is a POST to api/dives")
}
