namespace go demo

service DemoService {
    Response Echo(1: Request req);
    oneway void Send(1: Request req);
}

struct Request {
    1: string msg
}

struct Response {
    1: i8 code
    2: string msg
}
