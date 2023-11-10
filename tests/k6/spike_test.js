/* 
    Spike test is a variation of a stress test, but it does not gradually increase the load, instead it spikes to exterme load over short window of time

    Run a stress test to:
    - Determine how your system will perform under sudden surge of traffic
    - Determine if your system will recover once the traffic has subsided

    Success is based on expectation. Systems will generally react in 1 of 4 ways
    - Excellent: system performance is not degraded during the surge of trafic.
      Response time is similar during low traffic and high traffic
    - Good: Response time is slower, but the system does not produce any errors.
      All requests are handled
    - Poor: System produces errors during the surge of traffic, but recovers to normal after, the traffic subsides
    - Bad: System crashes, and does not recover after the traffic has subsided
*/

import http from "k6/http";

export let options = {
  insecureSkipTLSVerify: true,
  noConnectionReuse: false,
  stages: [
    { duration: "10s", target: 100 }, // below normal load
    { duration: "1m", target: 100 },
    { duration: "10s", target: 1400 }, // spike to 1400 users
    { duration: "3m", target: 1400 }, // stay at 1400 users for 3 minutes
    { duration: "10s", target: 100 }, // scale down, recovery stage.
    { duration: "3m", target: 100 }, // scale down, recovery stage.
    { duration: "10s", target: 0 }, // scale down, recovery stage.
  ],
};

export default () => {
  const url = "http://localhost:3001/twirp/twirp.v1.AuthService/CreateSession";
  const headers = { "Content-Type": "application/json" };
  const formData = JSON.stringify({ email: "hacero@gmail.com" });

  http.post(url, formData, { headers });
};
