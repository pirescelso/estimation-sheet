# Estimation Sheet - Test Results - REST API

### createManager
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Wed, 31 Jul 2024 16:22:02 GMT
Content-Length: 200
Connection: close

{
  "user_id": "ca0cd875-b18c-4be3-9e83-a55e15a62376",
  "email": "john.doe@userland.com",
  "user_name": "john1234",
  "name": "John Doe",
  "user_type": "manager",
  "created_at": "2024-07-31T16:22:02Z",
  "updated_at": null
}
```

### createEstimator
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Wed, 31 Jul 2024 16:24:19 GMT
Content-Length: 208
Connection: close

{
  "user_id": "057cd629-f097-485b-a1ca-2c728c1b4762",
  "email": "marie.doe1110@userland.com",
  "user_name": "marie123",
  "name": "Marie Doe",
  "user_type": "estimator",
  "created_at": "2024-07-31T16:24:19Z",
  "updated_at": null
}
```

### listUsers
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 31 Jul 2024 16:24:44 GMT
Content-Length: 420
Connection: close

{
  "users": [
    {
      "user_id": "ca0cd875-b18c-4be3-9e83-a55e15a62376",
      "email": "john.doe@userland.com",
      "user_name": "john1234",
      "name": "John Doe",
      "user_type": "manager",
      "created_at": "2024-07-31T16:22:02Z",
      "updated_at": null
    },
    {
      "user_id": "057cd629-f097-485b-a1ca-2c728c1b4762",
      "email": "marie.doe1110@userland.com",
      "user_name": "marie123",
      "name": "Marie Doe",
      "user_type": "estimator",
      "created_at": "2024-07-31T16:24:19Z",
      "updated_at": null
    }
  ]
}
```

### createBaseline
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Wed, 31 Jul 2024 16:25:10 GMT
Content-Length: 436
Connection: close

{
  "baseline_id": "015071c0-9249-4937-8da8-59b868712db7",
  "code": "RIT123456789",
  "review": 1,
  "title": "Logistics Cost \u0026 Time Management",
  "description": "This project will streamline our internal processes and increase overall efficiency",
  "duration": 12,
  "manager_id": "ca0cd875-b18c-4be3-9e83-a55e15a62376",
  "estimator_id": "057cd629-f097-485b-a1ca-2c728c1b4762",
  "start_date": "2024-01-01",
  "created_at": "2024-07-31T16:25:10Z",
  "updated_at": null
}
```

### listBaselines
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 31 Jul 2024 16:25:34 GMT
Content-Length: 497
Connection: close

{
  "baselines": [
    {
      "baseline_id": "015071c0-9249-4937-8da8-59b868712db7",
      "code": "RIT123456789",
      "review": 1,
      "title": "Logistics Cost \u0026 Time Management",
      "description": "This project will streamline our internal processes and increase overall efficiency",
      "duration": 12,
      "manager_id": "ca0cd875-b18c-4be3-9e83-a55e15a62376",
      "manager": "John Doe",
      "estimator_id": "057cd629-f097-485b-a1ca-2c728c1b4762",
      "estimator": "Marie Doe",
      "start_date": "2024-01-01",
      "created_at": "2024-07-31T16:25:10Z",
      "updated_at": null
    }
  ]
}
```

### createBaselineWithError
```json
HTTP/1.1 422 Unprocessable Entity
Content-Type: application/json
Date: Wed, 31 Jul 2024 16:25:59 GMT
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
Date: Wed, 31 Jul 2024 17:15:06 GMT
Content-Length: 454
Connection: close

{
  "baseline_id": "015071c0-9249-4937-8da8-59b868712db7",
  "code": "RIT123456789",
  "review": 1,
  "title": "Logistics Cost \u0026 Time Management",
  "description": "This project will streamline our internal processes and increase overall efficiency",
  "duration": 12,
  "manager_id": "ca0cd875-b18c-4be3-9e83-a55e15a62376",
  "estimator_id": "057cd629-f097-485b-a1ca-2c728c1b4762",
  "start_date": "2024-01-01",
  "created_at": "2024-07-31T16:25:10Z",
  "updated_at": "2024-07-31T17:15:06Z"
}
```

### getBaseline
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 31 Jul 2024 17:16:05 GMT
Content-Length: 499
Connection: close

{
  "baseline_id": "015071c0-9249-4937-8da8-59b868712db7",
  "code": "RIT123456789",
  "review": 1,
  "title": "Logistics Cost \u0026 Time Management",
  "description": "This project will streamline our internal processes and increase overall efficiency",
  "duration": 12,
  "manager_id": "ca0cd875-b18c-4be3-9e83-a55e15a62376",
  "manager": "John Doe",
  "estimator_id": "057cd629-f097-485b-a1ca-2c728c1b4762",
  "estimator": "Marie Doe",
  "start_date": "2024-01-01",
  "created_at": "2024-07-31T16:25:10Z",
  "updated_at": "2024-07-31T17:15:06Z"
}
```

### createCostPO
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Wed, 31 Jul 2024 17:16:26 GMT
Content-Length: 541
Connection: close

{
  "cost_id": "4b5539d1-186a-48db-94b8-c0decdcb6ded",
  "baseline_id": "015071c0-9249-4937-8da8-59b868712db7",
  "cost_type": "one_time",
  "description": "Mão de obra do PO",
  "comment": "estimativa do PO",
  "amount": 180000,
  "currency": "BRL",
  "tax": 0,
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
  "created_at": "2024-07-31T17:16:26Z",
  "updated_at": null
}
```

### createCostConsulting
```json
Content-Type: application/json
Date: Wed, 31 Jul 2024 17:17:10 GMT
Content-Length: 405
Connection: close

{
  "cost_id": "4322ebe2-316d-4bdc-8b02-55da075683cf",
  "baseline_id": "015071c0-9249-4937-8da8-59b868712db7",
  "cost_type": "one_time",
  "description": "External Consulting",
  "comment": "estimativa de consultoria externa",
  "amount": 80000,
  "currency": "EUR",
  "tax": 23.1,
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
  "created_at": "2024-07-31T17:17:10Z",
  "updated_at": null
}
```

### getCostsByBaselineId
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 31 Jul 2024 17:18:00 GMT
Content-Length: 958
Connection: close

{
  "costs": [
    {
      "cost_id": "4322ebe2-316d-4bdc-8b02-55da075683cf",
      "baseline_id": "015071c0-9249-4937-8da8-59b868712db7",
      "cost_type": "one_time",
      "description": "External Consulting",
      "comment": "estimativa de consultoria externa",
      "amount": 80000,
      "currency": "EUR",
      "tax": 23.1,
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
      "created_at": "2024-07-31T17:17:10Z",
      "updated_at": null
    },
    {
      "cost_id": "4b5539d1-186a-48db-94b8-c0decdcb6ded",
      "baseline_id": "015071c0-9249-4937-8da8-59b868712db7",
      "cost_type": "one_time",
      "description": "Mão de obra do PO",
      "comment": "estimativa do PO",
      "amount": 180000,
      "currency": "BRL",
      "tax": 0,
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
      "created_at": "2024-07-31T17:16:26Z",
      "updated_at": null
    }
  ]
}
```
### updateCostConsulting
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 31 Jul 2024 17:18:33 GMT
Content-Length: 432
Connection: close

{
  "cost_id": "4322ebe2-316d-4bdc-8b02-55da075683cf",
  "baseline_id": "015071c0-9249-4937-8da8-59b868712db7",
  "cost_type": "one_time",
  "description": "External Consulting",
  "comment": "estimativa de consultoria externa atualizada",
  "amount": 80000,
  "currency": "USD",
  "tax": 23,
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
  "created_at": "2024-07-31T17:17:10Z",
  "updated_at": "2024-07-31T17:18:33Z"
}
```

### createPlanBP
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Wed, 31 Jul 2024 17:19:48 GMT
Content-Length: 617
Connection: close

{
  "plan_id": "047e4b3e-e253-415e-a6c8-5eacd462bf3b",
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
  "created_at": "2024-07-31T17:19:48Z",
  "updated_at": null
}
```
### createPlanFC03
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Wed, 31 Jul 2024 17:20:28 GMT
Content-Length: 618
Connection: close

{
  "plan_id": "52cda920-e8d7-4d3a-a836-eba11c97e785",
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
  "created_at": "2024-07-31T17:20:28Z",
  "updated_at": null
}
```
### listPlans
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 31 Jul 2024 17:21:12 GMT
Content-Length: 313
Connection: close

{
  "plans": [
    {
      "plan_id": "047e4b3e-e253-415e-a6c8-5eacd462bf3b",
      "code": "BP 2025",
      "name": "Business Plan 2025",
      "created_at": "2024-07-31T17:19:48Z",
      "updated_at": null
    },
    {
      "plan_id": "52cda920-e8d7-4d3a-a836-eba11c97e785",
      "code": "FC 03 2025",
      "name": "Forecast 03 2025",
      "created_at": "2024-07-31T17:20:28Z",
      "updated_at": null
    }
  ]
}
```
### updatePlan
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 31 Jul 2024 17:21:44 GMT
Content-Length: 637
Connection: close

{
  "plan_id": "047e4b3e-e253-415e-a6c8-5eacd462bf3b",
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
  "created_at": "2024-07-31T17:19:48Z",
  "updated_at": "2024-07-31T17:21:44Z"
}
```

### createPortfolioBP
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Wed, 31 Jul 2024 17:22:26 GMT
Content-Length: 56
Connection: close

{
  "portfolio_id": "96d7f238-2371-4645-b4a6-77207bc9633c"
}
```
### createPortfolioFC03
```json
HTTP/1.1 201 Created
Content-Type: application/json
Date: Wed, 31 Jul 2024 17:22:49 GMT
Content-Length: 56
Connection: close

{
  "portfolio_id": "b36d4107-dd26-4b78-abb0-0934a059a278"
}
```

### listPortfoliosByPlan
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 31 Jul 2024 17:23:12 GMT
Content-Length: 418
Connection: close

{
  "portfolios": [
    {
      "portfolio_id": "b36d4107-dd26-4b78-abb0-0934a059a278",
      "code": "RIT123456789",
      "review": 1,
      "plan_code": "FC 03 2025",
      "title": "Logistics Cost \u0026 Time Management",
      "description": "This project will streamline our internal processes and increase overall efficiency",
      "duration": 12,
      "manager": "John Doe",
      "estimator": "Marie Doe",
      "start_date": "2025-07-01",
      "created_at": "2024-07-31T17:22:49Z",
      "updated_at": null
    }
  ]
}
```
### getPortfolioBP
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 31 Jul 2024 22:15:30 GMT
Content-Length: 1445
Connection: close

{
  "portfolio_id": "b36d4107-dd26-4b78-abb0-0934a059a278",
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
      "budget_id": "4f728851-b05a-495c-aff1-59b023a6f836",
      "portfolio_id": "b36d4107-dd26-4b78-abb0-0934a059a278",
      "cost_type": "one_time",
      "description": "External Consulting",
      "comment": "estimativa de consultoria externa atualizada",
      "cost_amount": 80000,
      "cost_currency": "USD",
      "cost_tax": 23,
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
      "created_at": "2024-07-31T17:22:49Z",
      "updated_at": null
    },
    {
      "budget_id": "d896fbf0-d860-4f05-ba60-ee3840108226",
      "portfolio_id": "b36d4107-dd26-4b78-abb0-0934a059a278",
      "cost_type": "one_time",
      "description": "Mão de obra do PO",
      "comment": "estimativa do PO",
      "cost_amount": 180000,
      "cost_currency": "BRL",
      "cost_tax": 0,
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
      "created_at": "2024-07-31T17:22:49Z",
      "updated_at": null
    }
  ],
  "start_date": "2025-07-01",
  "created_at": "2024-07-31T17:22:49Z",
  "updated_at": null
}
```
### getPortfolioFC03
```json
HTTP/1.1 200 OK
Content-Type: application/json
Date: Wed, 31 Jul 2024 22:22:08 GMT
Content-Length: 1445
Connection: close

{
  "portfolio_id": "b36d4107-dd26-4b78-abb0-0934a059a278",
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
      "budget_id": "4f728851-b05a-495c-aff1-59b023a6f836",
      "portfolio_id": "b36d4107-dd26-4b78-abb0-0934a059a278",
      "cost_type": "one_time",
      "description": "External Consulting",
      "comment": "estimativa de consultoria externa atualizada",
      "cost_amount": 80000,
      "cost_currency": "USD",
      "cost_tax": 23,
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
      "created_at": "2024-07-31T17:22:49Z",
      "updated_at": null
    },
    {
      "budget_id": "d896fbf0-d860-4f05-ba60-ee3840108226",
      "portfolio_id": "b36d4107-dd26-4b78-abb0-0934a059a278",
      "cost_type": "one_time",
      "description": "Mão de obra do PO",
      "comment": "estimativa do PO",
      "cost_amount": 180000,
      "cost_currency": "BRL",
      "cost_tax": 0,
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
      "created_at": "2024-07-31T17:22:49Z",
      "updated_at": null
    }
  ],
  "start_date": "2025-07-01",
  "created_at": "2024-07-31T17:22:49Z",
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
