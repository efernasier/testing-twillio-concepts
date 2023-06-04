# Twilio POC - Programmable voice and messaging.

## Overview

This repository contains a simple implementation of programmable calls and messaging offered as a part of the Twilio solutions.

This README aims to have a log of the experience for me while I was reading and exploring the Twilio solutions, including the goals, methodology, key findings, and lessons learned.

## Programmable voice and messaging goals

The primary objectives of this POC were as follows:

1. Integrates Twilio library in a Golang application and recognize what it involves.
2. Be able to send SMS and MMS by using Twilio Programmable messaging solution.
3. Be able to start a CALL by using Twilio Programmable voice solution.

## Methodology

The POC was conducted in the following steps:

1. **Planning and Research**: A Twilio API documentation review and exploration to know about the ecosystem and what tools are required to use their products.

2. **Environment Setup**: A dedicated environment was set up to conduct the POC, including the installation and configuration of the necessary golang libraries, open a twilio account, phone numbers creations and trial account usages.

3. **Implementation**: A small-scale implementation was developed to demonstrate the functionalities and features of the technology. Calls, Messaging.

4. **Testing and Evaluation**: Simple testing was performed to assess how SMS and Phone calls can be generated, how to work with webhooks, and cloud functions, and react to replies made by a potential user.

5. **Analysis and Findings**: Chalenges and the findings were documented. The strengths, weaknesses, opportunities, and potential risks associated with the technology were identified and evaluated.

6. **Lessons Learned**: Based on the analysis, key insights and lessons learned were documented. These insights included recommendations, best practices, and considerations for further development and integration of the technology.

## Key Findings

The POC yielded the following key findings:

1. **Feature Suitability**: Integrates the twilio library in golang is quite easy. It doesn't requires super specials requirements. A trial and verified account is required to obtain Credentials and Tokens for interact with the Twilio ecosystem. 

2. **Performance and Scalability**: Twilio can work under a schema like functions and services. Since they are offering cloud solutions to think about scalability is managed by them. Also, provides the capability to have a sticky number to reply to the user by using the same phone number(even when several phone numbers are working together).

3. **Integration Complexity**: The sign-up process was not straightforward for me. During the regular phone number verification, the OTP codes that I received all of the times for me it was expires. I had to create a support ticket and perform the phone number verification by using my phone and also my email.

4. **Learning Curve**: The learning curve associated with adopting the Voice and Messaging concepts and terminology was manageable (eg. DTMF, IVR, CPS, AMD, SIP, VoIP). I was able to quickly grasp the fundamental concepts and gained proficiency through hands-on experimentation.

## Conclusion

The Twilio POC provided valuable insights and opened my mind by identifying several real-life cases for using SMS and phone calls. Additionally, if we plan to use this kind of technology, it is important to consider integrating request management mechanisms (e.g., rate limiters) and implementing controls to avoid additional costs.

Please, feel free to create PRs or Enhance this POC as you want.
