import http from "k6/http";

export let options = {
  insecureSkipTLSVerify: true,
  noConnectionReuse: false,
  vus: 10,
  duration: "10s",
};

export default () => {
  const url = "http://localhost:3001/twirp/twirp.v1.AuthService/CreateSession";
  const headers = { "Content-Type": "application/json" };
  const formData = JSON.stringify({ email: "hacero@gmail.com" });

  http.post(url, formData, { headers });
};
