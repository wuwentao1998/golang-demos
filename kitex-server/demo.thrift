namespace go demo

service Demo {
    Response echo(1: Request req)
}

struct Request {
    1: string msg
}

struct Response {
    1: i8 code
    2: string msg
}
