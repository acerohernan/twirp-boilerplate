/* 
    Stress Testing is a type of load testing used to determine the limits of the system.
    The purpose of this test is to verify the stability and reliability of the system under exterme conditions

    Run a stress test to:
    - Determine how your system will behave under extreme condititions
    - Determine what is the maximun capacity of your system in terms of users of throught
    - Determine the breaking point of our system and its failure mode
    - Determine if your system will recover without manual intervention after the stress test is over

    More of a load test than a spike test
*/

import http from "k6/http";

export let options = {
  insecureSkipTLSVerify: true,
  noConnectionReuse: false,
  stages: [
    { duration: "2m", target: 100 }, // below normal load
    { duration: "5m", target: 100 },
    { duration: "2m", target: 200 }, // normal load
    { duration: "5m", target: 200 },
    { duration: "2m", target: 300 }, // around breaking point
    { duration: "5m", target: 300 },
    { duration: "5m", target: 400 }, // beyond breaking point
    { duration: "5m", target: 400 },
    { duration: "10m", target: 0 }, // scale down. recovery stage.
  ],
};

export default () => {
  const url = "http://localhost:3001/twirp/twirp.v1.AuthService/CreateSession";
  const headers = { "Content-Type": "application/json" };
  const formData = JSON.stringify({ email: "hacero@gmail.com" });

  http.post(url, formData, { headers });
};
