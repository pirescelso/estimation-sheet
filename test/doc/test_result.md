# Estimation Sheet - Test Results - REST API

### createManager
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Sun, 11 Aug 2024 15:34:08 GMT
Content-Length: 200
Connection: close

{
  "user_id": "80126c9b-c093-4722-b019-a8c5b7d6445c",
  "email": "john.doe@userland.com",
  "user_name": "john1234",
  "name": "John Doe",
  "user_type": "manager",
  "created_at": "2024-08-11T15:34:08Z",
  "updated_at": null
}
```

### createEstimator
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Sun, 11 Aug 2024 15:34:30 GMT
Content-Length: 208
Connection: close

{
  "user_id": "adbfa1e8-0443-49ab-b05b-ae6b1e7c917c",
  "email": "marie.doe1110@userland.com",
  "user_name": "marie123",
  "name": "Marie Doe",
  "user_type": "estimator",
  "created_at": "2024-08-11T15:34:30Z",
  "updated_at": null
}
```

### listUsers
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 11 Aug 2024 15:35:01 GMT
Content-Length: 420
Connection: close

{
  "users": [
    {
      "user_id": "80126c9b-c093-4722-b019-a8c5b7d6445c",
      "email": "john.doe@userland.com",
      "user_name": "john1234",
      "name": "John Doe",
      "user_type": "manager",
      "created_at": "2024-08-11T15:34:08Z",
      "updated_at": null
    },
    {
      "user_id": "adbfa1e8-0443-49ab-b05b-ae6b1e7c917c",
      "email": "marie.doe1110@userland.com",
      "user_name": "marie123",
      "name": "Marie Doe",
      "user_type": "estimator",
      "created_at": "2024-08-11T15:34:30Z",
      "updated_at": null
    }
  ]
}
```
### createCompetence
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Sun, 11 Aug 2024 15:36:07 GMT
Content-Length: 162
Connection: close

{
  "competence_id": "f0823918-4fc5-4794-a3b4-1530d1f5198d",
  "code": "Tech Doc",
  "name": "Technical Documentation",
  "created_at": "2024-08-11T15:36:07Z",
  "updated_at": null
}
```

### createBaseline
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Sun, 11 Aug 2024 15:37:11 GMT
Content-Length: 436
Connection: close

{
  "baseline_id": "bc79b3f0-71c9-432c-929f-465837ee24d8",
  "code": "RIT123456789",
  "review": 1,
  "title": "Logistics Cost \u0026 Time Management",
  "description": "This project will streamline our internal processes and increase overall efficiency",
  "duration": 12,
  "manager_id": "80126c9b-c093-4722-b019-a8c5b7d6445c",
  "estimator_id": "adbfa1e8-0443-49ab-b05b-ae6b1e7c917c",
  "start_date": "2024-01-01",
  "created_at": "2024-08-11T15:37:11Z",
  "updated_at": null
}
```

### listBaselines
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 11 Aug 2024 15:37:45 GMT
Content-Length: 497
Connection: close

{
  "baselines": [
    {
      "baseline_id": "bc79b3f0-71c9-432c-929f-465837ee24d8",
      "code": "RIT123456789",
      "review": 1,
      "title": "Logistics Cost \u0026 Time Management",
      "description": "This project will streamline our internal processes and increase overall efficiency",
      "duration": 12,
      "manager_id": "80126c9b-c093-4722-b019-a8c5b7d6445c",
      "manager": "John Doe",
      "estimator_id": "adbfa1e8-0443-49ab-b05b-ae6b1e7c917c",
      "estimator": "Marie Doe",
      "start_date": "2024-01-01",
      "created_at": "2024-08-11T15:37:11Z",
      "updated_at": null
    }
  ]
}
```

### createBaselineWithError
```json
HTTP/1.1 422 Unprocessable Entity
Content-Type: application/json
Date: Sun, 11 Aug 2024 15:38:14 GMT
Content-Length: 161
Connection: close

{
  "status_code": 422,
  "error": "Unprocessable Entity",
  "message": {
    "invalid_payload": [
      {
        "field": "estimator_id",
        "error": "EstimatorID must be a valid version 4 UUID"
      }
    ]
  }
}
```

### updateBaseline
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 11 Aug 2024 15:38:34 GMT
Content-Length: 436
Connection: close

{
  "baseline_id": "bc79b3f0-71c9-432c-929f-465837ee24d8",
  "code": "RIT123456789",
  "review": 1,
  "title": "Logistics Cost \u0026 Time Management",
  "description": "This project will streamline our internal processes and increase overall efficiency",
  "duration": 12,
  "manager_id": "80126c9b-c093-4722-b019-a8c5b7d6445c",
  "estimator_id": "adbfa1e8-0443-49ab-b05b-ae6b1e7c917c",
  "start_date": "2024-01-01",
  "created_at": "2024-08-11T15:37:11Z",
  "updated_at": null
}
```

### getBaseline
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 11 Aug 2024 15:38:54 GMT
Content-Length: 499
Connection: close

{
  "baseline_id": "bc79b3f0-71c9-432c-929f-465837ee24d8",
  "code": "RIT123456789",
  "review": 1,
  "title": "Logistics Cost \u0026 Time Management",
  "description": "This project will streamline our internal processes and increase overall efficiency",
  "duration": 12,
  "manager_id": "80126c9b-c093-4722-b019-a8c5b7d6445c",
  "manager": "John Doe",
  "estimator_id": "adbfa1e8-0443-49ab-b05b-ae6b1e7c917c",
  "estimator": "Marie Doe",
  "start_date": "2024-01-01",
  "created_at": "2024-08-11T15:37:11Z",
  "updated_at": "2024-08-11T15:38:34Z"
}
```

### createCostPO
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Sun, 11 Aug 2024 15:39:24 GMT
Content-Length: 541
Connection: close

{
  "cost_id": "d962f363-781a-4d73-a25a-b89a8b3a53d9",
  "baseline_id": "bc79b3f0-71c9-432c-929f-465837ee24d8",
  "cost_type": "one_time",
  "description": "Mão de obra do PO",
  "comment": "estimativa do PO",
  "amount": 180000,
  "currency": "BRL",
  "tax": 0,
  "apply_inflation": true,
  "cost_allocations": [
    {
      "year": 2024,
      "month": 1,
      "amount": 30000
    },
    {
      "year": 2024,
      "month": 2,
      "amount": 30000
    },
    {
      "year": 2024,
      "month": 3,
      "amount": 30000
    },
    {
      "year": 2024,
      "month": 4,
      "amount": 30000
    },
    {
      "year": 2024,
      "month": 5,
      "amount": 30000
    },
    {
      "year": 2024,
      "month": 6,
      "amount": 30000
    }
  ],
  "created_at": "2024-08-11T15:39:24Z",
  "updated_at": null
}
```

### createCostConsulting
```json
Content-Type: application/json
Date: Sun, 11 Aug 2024 15:39:51 GMT
Content-Length: 405
Connection: close

{
  "cost_id": "25967310-c4bd-42b3-bca6-f51a7a453c27",
  "baseline_id": "bc79b3f0-71c9-432c-929f-465837ee24d8",
  "cost_type": "one_time",
  "description": "External Consulting",
  "comment": "estimativa de consultoria externa",
  "amount": 80000,
  "currency": "EUR",
  "tax": 23.1,
  "apply_inflation": false,
  "cost_allocations": [
    {
      "year": 2024,
      "month": 4,
      "amount": 30000
    },
    {
      "year": 2024,
      "month": 6,
      "amount": 50000
    }
  ],
  "created_at": "2024-08-11T15:39:51Z",
  "updated_at": null
}
```

### getCostsByBaselineId
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 11 Aug 2024 15:40:13 GMT
Content-Length: 958
Connection: close

{
  "costs": [
    {
      "cost_id": "25967310-c4bd-42b3-bca6-f51a7a453c27",
      "baseline_id": "bc79b3f0-71c9-432c-929f-465837ee24d8",
      "cost_type": "one_time",
      "description": "External Consulting",
      "comment": "estimativa de consultoria externa",
      "amount": 80000,
      "currency": "EUR",
      "tax": 23.1,
      "apply_inflation": false,
      "cost_allocations": [
        {
          "year": 2024,
          "month": 4,
          "amount": 30000
        },
        {
          "year": 2024,
          "month": 6,
          "amount": 50000
        }
      ],
      "created_at": "2024-08-11T15:39:51Z",
      "updated_at": null
    },
    {
      "cost_id": "d962f363-781a-4d73-a25a-b89a8b3a53d9",
      "baseline_id": "bc79b3f0-71c9-432c-929f-465837ee24d8",
      "cost_type": "one_time",
      "description": "Mão de obra do PO",
      "comment": "estimativa do PO",
      "amount": 180000,
      "currency": "BRL",
      "tax": 0,
      "apply_inflation": true,
      "cost_allocations": [
        {
          "year": 2024,
          "month": 1,
          "amount": 30000
        },
        {
          "year": 2024,
          "month": 2,
          "amount": 30000
        },
        {
          "year": 2024,
          "month": 3,
          "amount": 30000
        },
        {
          "year": 2024,
          "month": 4,
          "amount": 30000
        },
        {
          "year": 2024,
          "month": 5,
          "amount": 30000
        },
        {
          "year": 2024,
          "month": 6,
          "amount": 30000
        }
      ],
      "created_at": "2024-08-11T15:39:24Z",
      "updated_at": null
    }
  ]
}
```
### updateCostConsulting
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 11 Aug 2024 15:40:42 GMT
Content-Length: 438
Connection: close

{
  "cost_id": "25967310-c4bd-42b3-bca6-f51a7a453c27",
  "baseline_id": "bc79b3f0-71c9-432c-929f-465837ee24d8",
  "cost_type": "one_time",
  "description": "External Consulting",
  "comment": "estimativa de consultoria externa atualizada",
  "amount": 80000,
  "currency": "USD",
  "tax": 23,
  "apply_inflation": false,
  "cost_allocations": [
    {
      "year": 2024,
      "month": 7,
      "amount": 50000
    },
    {
      "year": 2024,
      "month": 8,
      "amount": 30000
    }
  ],
  "created_at": "2024-08-11T15:39:51Z",
  "updated_at": null
}
```
### createEffort
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Sun, 11 Aug 2024 15:42:05 GMT
Content-Length: 576
Connection: close

{
  "effort_id": "bdb2599d-ecc1-4a96-a24e-b2c746ba3706",
  "baseline_id": "bc79b3f0-71c9-432c-929f-465837ee24d8",
  "competence_id": "f0823918-4fc5-4794-a3b4-1530d1f5198d",
  "comment": "considerado a Simone na atividade",
  "hours": 160,
  "effort_allocations": [
    {
      "year": 2024,
      "month": 1,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 2,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 3,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 4,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 5,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 6,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 7,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 8,
      "hours": 20
    }
  ],
  "created_at": "2024-08-11T15:42:05Z",
  "updated_at": null
}
```
### updateEffort
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 11 Aug 2024 15:43:13 GMT
Content-Length: 601
Connection: close

{
  "effort_id": "bdb2599d-ecc1-4a96-a24e-b2c746ba3706",
  "baseline_id": "bc79b3f0-71c9-432c-929f-465837ee24d8",
  "competence_id": "f0823918-4fc5-4794-a3b4-1530d1f5198d",
  "comment": "considerado Simone e Regina na atividade",
  "hours": 160,
  "effort_allocations": [
    {
      "year": 2024,
      "month": 1,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 2,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 3,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 4,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 5,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 6,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 7,
      "hours": 20
    },
    {
      "year": 2024,
      "month": 8,
      "hours": 20
    }
  ],
  "created_at": "2024-08-11T15:42:05Z",
  "updated_at": "2024-08-11T15:43:13Z"
}
```
### listEfforts
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 11 Aug 2024 15:44:12 GMT
Content-Length: 615
Connection: close

{
  "efforts": [
    {
      "effort_id": "bdb2599d-ecc1-4a96-a24e-b2c746ba3706",
      "baseline_id": "bc79b3f0-71c9-432c-929f-465837ee24d8",
      "competence_id": "f0823918-4fc5-4794-a3b4-1530d1f5198d",
      "comment": "considerado Simone e Regina na atividade",
      "hours": 160,
      "effort_allocations": [
        {
          "year": 2024,
          "month": 1,
          "hours": 20
        },
        {
          "year": 2024,
          "month": 2,
          "hours": 20
        },
        {
          "year": 2024,
          "month": 3,
          "hours": 20
        },
        {
          "year": 2024,
          "month": 4,
          "hours": 20
        },
        {
          "year": 2024,
          "month": 5,
          "hours": 20
        },
        {
          "year": 2024,
          "month": 6,
          "hours": 20
        },
        {
          "year": 2024,
          "month": 7,
          "hours": 20
        },
        {
          "year": 2024,
          "month": 8,
          "hours": 20
        }
      ],
      "created_at": "2024-08-11T15:42:05Z",
      "updated_at": "2024-08-11T15:43:13Z"
    }
  ]
}
```

### createPlanBP
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Sun, 11 Aug 2024 15:44:39 GMT
Content-Length: 617
Connection: close

{
  "plan_id": "e859fcdb-6322-4714-80ae-c263f38a645c",
  "code": "BP 2025",
  "name": "Business Plan 2025",
  "assumptions": [
    {
      "year": 2024,
      "inflation": 4,
      "currencies": [
        {
          "currency": "USD",
          "exchange": 4.5
        },
        {
          "currency": "EUR",
          "exchange": 5.5
        }
      ]
    },
    {
      "year": 2025,
      "inflation": 5.2,
      "currencies": [
        {
          "currency": "USD",
          "exchange": 5
        },
        {
          "currency": "EUR",
          "exchange": 6
        }
      ]
    },
    {
      "year": 2026,
      "inflation": 5.26,
      "currencies": [
        {
          "currency": "USD",
          "exchange": 5.55
        },
        {
          "currency": "EUR",
          "exchange": 6.66
        }
      ]
    },
    {
      "year": 2027,
      "inflation": 5.3,
      "currencies": [
        {
          "currency": "USD",
          "exchange": 5.77
        },
        {
          "currency": "EUR",
          "exchange": 6.88
        }
      ]
    }
  ],
  "created_at": "2024-08-11T15:44:39Z",
  "updated_at": null
}
```
### createPlanFC03
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Sun, 11 Aug 2024 15:45:09 GMT
Content-Length: 618
Connection: close

{
  "plan_id": "b1ede63d-ead4-4c7c-aaa5-cc0edd35dfe2",
  "code": "FC 03 2025",
  "name": "Forecast 03 2025",
  "assumptions": [
    {
      "year": 2024,
      "inflation": 4,
      "currencies": [
        {
          "currency": "USD",
          "exchange": 4.5
        },
        {
          "currency": "EUR",
          "exchange": 5.5
        }
      ]
    },
    {
      "year": 2025,
      "inflation": 5.2,
      "currencies": [
        {
          "currency": "USD",
          "exchange": 5
        },
        {
          "currency": "EUR",
          "exchange": 6
        }
      ]
    },
    {
      "year": 2026,
      "inflation": 5.26,
      "currencies": [
        {
          "currency": "USD",
          "exchange": 5.55
        },
        {
          "currency": "EUR",
          "exchange": 6.66
        }
      ]
    },
    {
      "year": 2027,
      "inflation": 5.3,
      "currencies": [
        {
          "currency": "USD",
          "exchange": 5.77
        },
        {
          "currency": "EUR",
          "exchange": 6.88
        }
      ]
    }
  ],
  "created_at": "2024-08-11T15:45:09Z",
  "updated_at": null
}
```
### listPlans
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 11 Aug 2024 15:46:17 GMT
Content-Length: 1247
Connection: close

{
  "plans": [
    {
      "plan_id": "e859fcdb-6322-4714-80ae-c263f38a645c",
      "code": "BP 2025",
      "name": "Business Plan 2025",
      "assumptions": [
        {
          "year": 2024,
          "inflation": 4,
          "currencies": [
            {
              "currency": "USD",
              "exchange": 4.5
            },
            {
              "currency": "EUR",
              "exchange": 5.5
            }
          ]
        },
        {
          "year": 2025,
          "inflation": 5.2,
          "currencies": [
            {
              "currency": "USD",
              "exchange": 5
            },
            {
              "currency": "EUR",
              "exchange": 6
            }
          ]
        },
        {
          "year": 2026,
          "inflation": 5.26,
          "currencies": [
            {
              "currency": "USD",
              "exchange": 5.55
            },
            {
              "currency": "EUR",
              "exchange": 6.66
            }
          ]
        },
        {
          "year": 2027,
          "inflation": 5.3,
          "currencies": [
            {
              "currency": "USD",
              "exchange": 5.77
            },
            {
              "currency": "EUR",
              "exchange": 6.88
            }
          ]
        }
      ],
      "created_at": "2024-08-11T15:44:39Z",
      "updated_at": null
    },
    {
      "plan_id": "b1ede63d-ead4-4c7c-aaa5-cc0edd35dfe2",
      "code": "FC 03 2025",
      "name": "Forecast 03 2025",
      "assumptions": [
        {
          "year": 2024,
          "inflation": 4,
          "currencies": [
            {
              "currency": "USD",
              "exchange": 4.5
            },
            {
              "currency": "EUR",
              "exchange": 5.5
            }
          ]
        },
        {
          "year": 2025,
          "inflation": 5.2,
          "currencies": [
            {
              "currency": "USD",
              "exchange": 5
            },
            {
              "currency": "EUR",
              "exchange": 6
            }
          ]
        },
        {
          "year": 2026,
          "inflation": 5.26,
          "currencies": [
            {
              "currency": "USD",
              "exchange": 5.55
            },
            {
              "currency": "EUR",
              "exchange": 6.66
            }
          ]
        },
        {
          "year": 2027,
          "inflation": 5.3,
          "currencies": [
            {
              "currency": "USD",
              "exchange": 5.77
            },
            {
              "currency": "EUR",
              "exchange": 6.88
            }
          ]
        }
      ],
      "created_at": "2024-08-11T15:45:09Z",
      "updated_at": null
    }
  ]
}
```
### updatePlan
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 11 Aug 2024 15:47:34 GMT
Content-Length: 637
Connection: close

{
  "plan_id": "e859fcdb-6322-4714-80ae-c263f38a645c",
  "code": "BP 2025",
  "name": "Business Plan 2025",
  "assumptions": [
    {
      "year": 2024,
      "inflation": 4,
      "currencies": [
        {
          "currency": "USD",
          "exchange": 4.55
        },
        {
          "currency": "EUR",
          "exchange": 5.55
        }
      ]
    },
    {
      "year": 2025,
      "inflation": 5.2,
      "currencies": [
        {
          "currency": "USD",
          "exchange": 5
        },
        {
          "currency": "EUR",
          "exchange": 6
        }
      ]
    },
    {
      "year": 2026,
      "inflation": 5.26,
      "currencies": [
        {
          "currency": "USD",
          "exchange": 5.55
        },
        {
          "currency": "EUR",
          "exchange": 6.66
        }
      ]
    },
    {
      "year": 2027,
      "inflation": 5.3,
      "currencies": [
        {
          "currency": "USD",
          "exchange": 5.77
        },
        {
          "currency": "EUR",
          "exchange": 6.88
        }
      ]
    }
  ],
  "created_at": "2024-08-11T15:44:39Z",
  "updated_at": "2024-08-11T15:47:34Z"
}
```

### createPortfolioBP
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Sun, 11 Aug 2024 15:48:17 GMT
Content-Length: 56
Connection: close

{
  "portfolio_id": "2c798462-c0f1-484d-ad41-13cbcabc914d"
}
```
### createPortfolioFC03
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Sun, 11 Aug 2024 15:48:36 GMT
Content-Length: 56
Connection: close

{
  "portfolio_id": "ca37af01-47cd-44a1-8b2b-6590c79f298e"
}
```

### listPortfoliosFC03
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 11 Aug 2024 15:49:58 GMT
Content-Length: 418
Connection: close

{
  "portfolios": [
    {
      "portfolio_id": "ca37af01-47cd-44a1-8b2b-6590c79f298e",
      "code": "RIT123456789",
      "review": 1,
      "plan_code": "FC 03 2025",
      "title": "Logistics Cost \u0026 Time Management",
      "description": "This project will streamline our internal processes and increase overall efficiency",
      "duration": 12,
      "manager": "John Doe",
      "estimator": "Marie Doe",
      "start_date": "2025-07-01",
      "created_at": "2024-08-11T15:48:36Z",
      "updated_at": null
    }
  ]
}
```
### listPortfoliosBP
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 11 Aug 2024 15:52:00 GMT
Content-Length: 415
Connection: close

{
  "portfolios": [
    {
      "portfolio_id": "2c798462-c0f1-484d-ad41-13cbcabc914d",
      "code": "RIT123456789",
      "review": 1,
      "plan_code": "BP 2025",
      "title": "Logistics Cost \u0026 Time Management",
      "description": "This project will streamline our internal processes and increase overall efficiency",
      "duration": 12,
      "manager": "John Doe",
      "estimator": "Marie Doe",
      "start_date": "2024-12-01",
      "created_at": "2024-08-11T15:48:17Z",
      "updated_at": null
    }
  ]
}
```

### getPortfolioBP
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 11 Aug 2024 15:52:31 GMT
Connection: close
Transfer-Encoding: chunked

{
  "portfolio_id": "2c798462-c0f1-484d-ad41-13cbcabc914d",
  "code": "RIT123456789",
  "review": 1,
  "plan_code": "BP 2025",
  "title": "Logistics Cost \u0026 Time Management",
  "description": "This project will streamline our internal processes and increase overall efficiency",
  "duration": 12,
  "manager": "John Doe",
  "estimator": "Marie Doe",
  "budgets": [
    {
      "budget_id": "f61ea160-e170-45e2-abe3-368edcb47964",
      "portfolio_id": "2c798462-c0f1-484d-ad41-13cbcabc914d",
      "cost_type": "one_time",
      "description": "External Consulting",
      "comment": "estimativa de consultoria externa atualizada",
      "cost_amount": 80000,
      "cost_currency": "USD",
      "cost_tax": 23,
      "cost_apply_inflation": false,
      "amount": 492000,
      "budget_allocations": [
        {
          "year": 2025,
          "month": 6,
          "amount": 307500
        },
        {
          "year": 2025,
          "month": 7,
          "amount": 184500
        }
      ],
      "created_at": "2024-08-11T15:48:17Z",
      "updated_at": null
    },
    {
      "budget_id": "1ff79543-51a2-4da5-bffd-d4acaed9d445",
      "portfolio_id": "2c798462-c0f1-484d-ad41-13cbcabc914d",
      "cost_type": "one_time",
      "description": "Mão de obra do PO",
      "comment": "estimativa do PO",
      "cost_amount": 180000,
      "cost_currency": "BRL",
      "cost_tax": 0,
      "cost_apply_inflation": true,
      "amount": 187800,
      "budget_allocations": [
        {
          "year": 2024,
          "month": 12,
          "amount": 30000
        },
        {
          "year": 2025,
          "month": 1,
          "amount": 31560
        },
        {
          "year": 2025,
          "month": 2,
          "amount": 31560
        },
        {
          "year": 2025,
          "month": 3,
          "amount": 31560
        },
        {
          "year": 2025,
          "month": 4,
          "amount": 31560
        },
        {
          "year": 2025,
          "month": 5,
          "amount": 31560
        }
      ],
      "created_at": "2024-08-11T15:48:17Z",
      "updated_at": null
    }
  ],
  "workloads": [
    {
      "workload_id": "01d4681c-29e6-43c0-b22d-900580651666",
      "portfolio_id": "2c798462-c0f1-484d-ad41-13cbcabc914d",
      "competence_code": "Tech Doc",
      "competence_name": "Technical Documentation",
      "comment": "considerado Simone e Regina na atividade",
      "hours": 160,
      "workload_allocations": [
        {
          "year": 2024,
          "month": 12,
          "hours": 20
        },
        {
          "year": 2025,
          "month": 1,
          "hours": 20
        },
        {
          "year": 2025,
          "month": 2,
          "hours": 20
        },
        {
          "year": 2025,
          "month": 3,
          "hours": 20
        },
        {
          "year": 2025,
          "month": 4,
          "hours": 20
        },
        {
          "year": 2025,
          "month": 5,
          "hours": 20
        },
        {
          "year": 2025,
          "month": 6,
          "hours": 20
        },
        {
          "year": 2025,
          "month": 7,
          "hours": 20
        }
      ],
      "created_at": "2024-08-11T15:48:17Z",
      "updated_at": null
    }
  ],
  "start_date": "2024-12-01",
  "created_at": "2024-08-11T15:48:17Z",
  "updated_at": null
}
```
### getPortfolioFC03
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Sun, 11 Aug 2024 15:53:16 GMT
Connection: close
Transfer-Encoding: chunked

{
  "portfolio_id": "ca37af01-47cd-44a1-8b2b-6590c79f298e",
  "code": "RIT123456789",
  "review": 1,
  "plan_code": "FC 03 2025",
  "title": "Logistics Cost \u0026 Time Management",
  "description": "This project will streamline our internal processes and increase overall efficiency",
  "duration": 12,
  "manager": "John Doe",
  "estimator": "Marie Doe",
  "budgets": [
    {
      "budget_id": "b9c2bbb4-0887-4b93-b8e4-b4b78012206b",
      "portfolio_id": "ca37af01-47cd-44a1-8b2b-6590c79f298e",
      "cost_type": "one_time",
      "description": "External Consulting",
      "comment": "estimativa de consultoria externa atualizada",
      "cost_amount": 80000,
      "cost_currency": "USD",
      "cost_tax": 23,
      "cost_apply_inflation": false,
      "amount": 546120,
      "budget_allocations": [
        {
          "year": 2026,
          "month": 1,
          "amount": 341325
        },
        {
          "year": 2026,
          "month": 2,
          "amount": 204795
        }
      ],
      "created_at": "2024-08-11T15:48:36Z",
      "updated_at": null
    },
    {
      "budget_id": "bf2ea9c5-294a-47d1-85ab-e6fb2fdb9b49",
      "portfolio_id": "ca37af01-47cd-44a1-8b2b-6590c79f298e",
      "cost_type": "one_time",
      "description": "Mão de obra do PO",
      "comment": "estimativa do PO",
      "cost_amount": 180000,
      "cost_currency": "BRL",
      "cost_tax": 0,
      "cost_apply_inflation": true,
      "amount": 189360,
      "budget_allocations": [
        {
          "year": 2025,
          "month": 7,
          "amount": 31560
        },
        {
          "year": 2025,
          "month": 8,
          "amount": 31560
        },
        {
          "year": 2025,
          "month": 9,
          "amount": 31560
        },
        {
          "year": 2025,
          "month": 10,
          "amount": 31560
        },
        {
          "year": 2025,
          "month": 11,
          "amount": 31560
        },
        {
          "year": 2025,
          "month": 12,
          "amount": 31560
        }
      ],
      "created_at": "2024-08-11T15:48:36Z",
      "updated_at": null
    }
  ],
  "workloads": [
    {
      "workload_id": "23ec5eee-fc75-4717-a8fb-ea103035269d",
      "portfolio_id": "ca37af01-47cd-44a1-8b2b-6590c79f298e",
      "competence_code": "Tech Doc",
      "competence_name": "Technical Documentation",
      "comment": "considerado Simone e Regina na atividade",
      "hours": 160,
      "workload_allocations": [
        {
          "year": 2025,
          "month": 7,
          "hours": 20
        },
        {
          "year": 2025,
          "month": 8,
          "hours": 20
        },
        {
          "year": 2025,
          "month": 9,
          "hours": 20
        },
        {
          "year": 2025,
          "month": 10,
          "hours": 20
        },
        {
          "year": 2025,
          "month": 11,
          "hours": 20
        },
        {
          "year": 2025,
          "month": 12,
          "hours": 20
        },
        {
          "year": 2026,
          "month": 1,
          "hours": 20
        },
        {
          "year": 2026,
          "month": 2,
          "hours": 20
        }
      ],
      "created_at": "2024-08-11T15:48:36Z",
      "updated_at": null
    }
  ],
  "start_date": "2025-07-01",
  "created_at": "2024-08-11T15:48:36Z",
  "updated_at": null
}
```
### deleteCostConsulting
```json
HTTP/1.1 409 Conflict
Content-Type: application/json
Date: Wed, 31 Jul 2024 22:23:23 GMT
Content-Length: 116
Connection: close

{
  "status_code": 409,
  "error": "Conflict",
  "message": "baseline 015071c0-9249-4937-8da8-59b868712db7 has 2 portfolio(s)"
}
```
### updateCostConsulting
```json
HTTP/1.1 409 Conflict
Content-Type: application/json
Date: Wed, 31 Jul 2024 22:25:08 GMT
Content-Length: 116
Connection: close

{
  "status_code": 409,
  "error": "Conflict",
  "message": "baseline 015071c0-9249-4937-8da8-59b868712db7 has 2 portfolio(s)"
}
```

### updateBaseline
```json
HTTP/1.1 409 Conflict
Content-Type: application/json
Date: Wed, 31 Jul 2024 22:27:38 GMT
Content-Length: 116
Connection: close

{
  "status_code": 409,
  "error": "Conflict",
  "message": "baseline 015071c0-9249-4937-8da8-59b868712db7 has 2 portfolio(s)"
}
```
### deletePortfolioBP
```json
HTTP/1.1 204 No Content
Content-Type: application/json
Date: Wed, 31 Jul 2024 22:27:57 GMT
Connection: close
```

### deletePortfolioFC03
```json
HTTP/1.1 204 No Content
Content-Type: application/json
Date: Wed, 31 Jul 2024 22:28:27 GMT
Connection: close
```
