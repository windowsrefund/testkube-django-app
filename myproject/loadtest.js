import http from "k6/http";
import { check } from "k6";

export default function() {
    let url = "http://django-test.com";

    let res = http.get(url);

    check(res, {
        "status is 200": (r) => r.status === 200,
        "text body contains welcome message": (r) => r.body.includes("The install worked successfully! Congratulations")
    });
}
