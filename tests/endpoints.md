# List of all endpoints
## Users 
```bash
POST http://localhost:9000/api/v1/users
PATCH http://localhost:9000/api/v1/users/{userID}
DELETE http://localhost:9000/api/v1/users/{userID}
GET http://localhost:9000/api/v1/users/{userID}
GET http://localhost:9000/api/v1/users
```
## Plans
```bash
POST http://localhost:9000/api/v1/plans
PATCH http://localhost:9000/api/v1/plans/{planID}
DELETE http://localhost:9000/api/v1/plans/{planID}
GET http://localhost:9000/api/v1/plans/{planID}
GET http://localhost:9000/api/v1/plans
```
## Baselines
```bash	
POST http://localhost:9000/api/baselines
PATCH http://localhost:9000/api/baselines/{baselineID}
DELETE http://localhost:9000/api/baselines/{baselineID}
GET http://localhost:9000/api/baselines/{baselineID}
GET http://localhost:9000/api/baselines
POST http://localhost:9000/api/baselines/{baselineID}/costs
PATCH http://localhost:9000/api/baselines/{baselineID}/costs/{costID}
DELETE http://localhost:9000/api/baselines/{baselineID}/costs/{costID}
GET http://localhost:9000/api/baselines/{baselineID}/costs
```
### Portfolios
```bash
POST POST http://localhost:9000/api/portfolios
DELETE http://localhost:9000/api/portfolios/{portfolioID}
GET http://localhost:9000/api/portfolios/{portfolioID}
GET http://localhost:9000/api/portfolios
GET http://localhost:9000/api/portfolios?planID={planID}
```