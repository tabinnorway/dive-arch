use actix_web::{get, post, HttpResponse, Responder, Result};
use actix_files::NamedFile;
use std::path::PathBuf;

pub fn get_file(fname: String) -> Result<NamedFile> {
    let path: String = format!("./public/{}", fname);
    let path: PathBuf = path.parse().unwrap();
    Ok(NamedFile::open(path)?)
}

#[get("/")]
pub async fn index() -> Result<NamedFile> {
    get_file("index.html".to_string())
}

#[get("/index.html")]
pub async fn index_file() -> Result<NamedFile> {
    get_file("index.html".to_string())
}

#[post("/echo")]
pub async fn echo(req_body: String) -> impl Responder {
    HttpResponse::Ok().body(req_body)
}

pub async fn manual_hello() -> impl Responder {
    HttpResponse::Ok().body("Hey there!")
}

pub async fn not_found() -> Result<NamedFile> {
    get_file("404.html".to_string())
}
