# GoCloudWatch

GoCloudWatch is a **serverless application** built with **Go**, leveraging **AWS Lambda, API Gateway, DynamoDB, and CloudWatch** for real-time monitoring, logging, and API management. The project focuses on **scalability, security, and performance optimization** in a cloud environment.

## Features 🚀

✅ **AWS CDK for Infrastructure as Code (IaC)**  
✅ **Serverless API powered by AWS Lambda & API Gateway**  
✅ **Secure authentication with JWT & bcrypt**  
✅ **CloudWatch Logs for monitoring API requests & performance**  
✅ **DynamoDB for scalable and serverless data storage**  
✅ **Middleware for protected API routes**  
✅ **API Testing with cURL commands**

---

## Project Setup 🛠️

### 1. Initialize the AWS CDK App

```sh
cdk init app --language go
```

### 2. Deploy Infrastructure

Check differences between local and deployed stack:

```sh
cdk diff
```

Deploy the stack:

```sh
cdk deploy
```

---

## AWS Services Used 🏗️

- **AWS Lambda** - Handles backend logic for the API.
- **API Gateway** - Exposes API endpoints securely.
- **DynamoDB** - Stores authentication and application data.
- **CloudWatch** - Logs and monitors API activity.
- **CloudFormation** - Manages AWS resources via Infrastructure as Code (IaC).

---

## API Endpoints 🌍

| Method | Endpoint     | Description                                            |
| ------ | ------------ | ------------------------------------------------------ |
| POST   | `/register`  | Registers a new user                                   |
| POST   | `/login`     | Authenticates a user and returns JWT token             |
| GET    | `/protected` | Protected route accessible only by authenticated users |

### Authentication Middleware 🔐

- Uses **JWT** for secure authentication.
- Passwords are hashed using **bcrypt**.
- Middleware ensures only authenticated users can access protected routes.

---

## Testing the API with cURL 🧪

### 1. Register a new user

```sh
curl -X POST https://your-api-gateway-url/register -d '{"email": "test@example.com", "password": "password123"}' -H "Content-Type: application/json"
```

### 2. Log in to get JWT token

```sh
curl -X POST https://your-api-gateway-url/login -d '{"email": "test@example.com", "password": "password123"}' -H "Content-Type: application/json"
```

### 3. Access protected route (Replace `TOKEN` with actual JWT)

```sh
curl -X GET https://your-api-gateway-url/protected -H "Authorization: Bearer TOKEN"
```

---

## Monitoring with CloudWatch 📊

View logs using AWS CLI:

```sh
aws logs tail /aws/lambda/my-lambda-function --follow
```

---

## Conclusion 🎯

GoCloudWatch demonstrates **scalable, serverless API design** using **Go** and **AWS services**. By integrating **CloudWatch Logs, API Gateway, and DynamoDB**, the project ensures **efficient monitoring, security, and performance**.

## Contribution Guidelines

Feel free to contribute, suggest improvements, or fork the project for your own cloud-based applications! 🚀

1. Fork the repository.
2. Create a new branch for your feature/bugfix.
3. Submit a pull request with a detailed description.

---

## License

This project is licensed under the [MIT License](LICENSE).

---

## Acknowledgments

- Inspired by [AWS CDK](https://aws.amazon.com/cdk/).
- Built with love using Go, AWS, and serverless architecture.

#HappyCoding 🎉
