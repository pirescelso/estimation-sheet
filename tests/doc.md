# Estimation Sheet - API

POST Users - manager 
```json
POST http://localhost:9000/api/v1/users
Content-Type: application/json

{
    "email": "john.doe@userland.com",
    "user_name": "john1234",
    "name": "John Doe",
    "user_type": "manager"
}

HTTP/1.1 201 Created
Content-Type: application/json
Date: Mon, 29 Jul 2024 00:42:10 GMT
Content-Length: 200
Connection: close

{
  "user_id": "238e207b-4900-4bbf-89a0-780a5b651c5d",
  "email": "john.doe@userland.com",
  "user_name": "john1234",
  "name": "John Doe",
  "user_type": "manager",
  "created_at": "2024-07-29T00:42:10Z",
  "updated_at": null
}

```
POST Users - estimator 
```json
POST http://localhost:9000/api/v1/users
Content-Type: application/json
{
    "email": "marie.doe1110@userland.com",
    "user_name": "marie123",
    "name": "Marie Doe",
    "user_type": "estimator"
}

HTTP/1.1 201 Created
Content-Type: application/json
Date: Mon, 29 Jul 2024 00:46:12 GMT
Content-Length: 208
Connection: close

{
  "user_id": "69e18d2a-0926-4498-9f9c-b7325747dbd6",
  "email": "marie.doe1110@userland.com",
  "user_name": "marie123",
  "name": "Marie Doe",
  "user_type": "estimator",
  "created_at": "2024-07-29T00:46:12Z",
  "updated_at": null
}
```
POST Baselines 
```json
POST http://localhost:9000/api/v1/baselines
Content-Type: application/json
{
    "code": "RIT123456789",
    "title": "Logistics Cost & Time Management",
    "description": "This project will streamline our internal processes and increase overall efficiency",
    "start_month": 1,
    "start_year": 2024,
    "duration": 12,
    "manager_id": "{{ managerId }}",
    "estimator_id": "{{ estimatorId }}"
}

HTTP/1.1 201 Created
Content-Type: application/json
Date: Mon, 29 Jul 2024 00:47:58 GMT
Content-Length: 425
Connection: close

{
  "baseline_id": "cb05970f-7e78-4915-84f1-62a898485f11",
  "code": "RIT123456789",
  "title": "Logistics Cost \u0026 Time Management",
  "description": "This project will streamline our internal processes and increase overall efficiency",
  "duration": 12,
  "manager_id": "238e207b-4900-4bbf-89a0-780a5b651c5d",
  "estimator_id": "69e18d2a-0926-4498-9f9c-b7325747dbd6",
  "start_date": "2024-01-01",
  "created_at": "2024-07-29T00:47:58Z",
  "updated_at": null
}
```
POST Costs - BRL
```json
POST http://localhost:9000/api/v1/costs
Content-Type: application/json
{
    "baseline_id": "{{ baselineId }}",
    "cost_type": "one_time",
    "description": "Mão de obra do PO",
    "comment": "estimativa do PO",
    "amount": 180000,
    "currency": "BRL",
    "tax": 0.00,
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
    ]
}

HTTP/1.1 201 Created
Content-Type: application/json
Date: Mon, 29 Jul 2024 00:50:13 GMT
Content-Length: 541
Connection: close

{
  "cost_id": "b5bf4e1a-d302-4963-b251-54039cf6b21d",
  "baseline_id": "cb05970f-7e78-4915-84f1-62a898485f11",
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
  "created_at": "2024-07-29T00:50:13Z",
  "updated_at": null
}
```
POST Cost - EUR
```json
POST http://localhost:9000/api/v1/costs
Content-Type: application/json
{
    "baseline_id": "{{ baselineId }}",
    "cost_type": "one_time",
    "description": "External Consulting",
    "comment": "estimativa de consultoria externa",
    "amount": 80000,
    "currency": "EUR",
    "tax": 23.10,
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
    ]
}

HTTP/1.1 201 Created
Content-Type: application/json
Date: Mon, 29 Jul 2024 00:51:59 GMT
Content-Length: 405
Connection: close

{
  "cost_id": "fbd01359-f977-4fa8-a058-822ac609ff77",
  "baseline_id": "cb05970f-7e78-4915-84f1-62a898485f11",
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
  "created_at": "2024-07-29T00:51:59Z",
  "updated_at": null
}
```
POST Plans -  BP 2025
```json
POST http://localhost:9000/api/v1/plans
Content-Type: application/json
{
    "code": "BP 2025",
    "name": "Business Plan 2025",
    "assumptions": [
        {
            "year": 2024,
            "inflation": 4.00,
            "currencies": [
                {
                    "currency": "USD",
                    "exchange": 4.50
                },
                {
                    "currency": "EUR",
                    "exchange": 5.50
                }
            ]
        },
        {
            "year": 2025,
            "inflation": 5.20,
            "currencies": [
                {
                    "currency": "USD",
                    "exchange": 5.00
                },
                {
                    "currency": "EUR",
                    "exchange": 6.00
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
            "inflation": 5.30,
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
    ]
}

HTTP/1.1 201 Created
Content-Type: application/json
Date: Mon, 29 Jul 2024 00:52:49 GMT
Content-Length: 617
Connection: close

{
  "plan_id": "dd0d6304-c945-40f8-beaf-b1e339bbebc5",
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
  "created_at": "2024-07-29T00:52:49Z",
  "updated_at": null
}
```
POST Plans - FC 03 2025
```json
POST http://localhost:9000/api/v1/plans
Content-Type: application/json
{
    "code": "FC 03 2025",
    "name": "Forecast 03 2025",
    "assumptions": [
        {
            "year": 2024,
            "inflation": 4.00,
            "currencies": [
                {
                    "currency": "USD",
                    "exchange": 4.50
                },
                {
                    "currency": "EUR",
                    "exchange": 5.50
                }
            ]
        },
        {
            "year": 2025,
            "inflation": 5.20,
            "currencies": [
                {
                    "currency": "USD",
                    "exchange": 5.00
                },
                {
                    "currency": "EUR",
                    "exchange": 6.00
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
            "inflation": 5.30,
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
    ]
}

HTTP/1.1 201 Created
Content-Type: application/json
Date: Mon, 29 Jul 2024 00:54:16 GMT
Content-Length: 618
Connection: close

{
  "plan_id": "1021af1e-4003-4f27-91ea-8feb231ca5d4",
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
  "created_at": "2024-07-29T00:54:16Z",
  "updated_at": null
}
```
POST Portfolios - BP 2025
```json
POST http://localhost:9000/api/v1/portfolios
Content-Type: application/json
{
    "baseline_id": "{{ baselineId }}",
    "plan_id": "{{ planIdBP }}",
    "shift_months": 11
}

HTTP/1.1 201 Created
Content-Type: application/json
Date: Mon, 29 Jul 2024 00:55:50 GMT
Content-Length: 56
Connection: close

{
  "portfolio_id": "0821c67e-d331-4c60-9257-75ed21590e44"
}
```
POST Portfolios - FC 03 2025
```json
POST http://localhost:9000/api/v1/portfolios
Content-Type: application/json
{
    "baseline_id": "{{ baselineId }}",
    "plan_id": "{{ planIdFC03 }}",
    "shift_months": 18
}

HTTP/1.1 201 Created
Content-Type: application/json
Date: Mon, 29 Jul 2024 00:57:32 GMT
Content-Length: 56
Connection: close

{
  "portfolio_id": "7670f871-e36a-4174-8888-b2f890442c7c"
}
```
GET Portfolios BP 2025
```json
GET http://localhost:9000/api/v1/portfolios/{{ portfolioIdBP }}

HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 29 Jul 2024 15:20:27 GMT
Content-Length: 1420
Connection: close

{
  "portfolio": {
    "portfolio_id": "0821c67e-d331-4c60-9257-75ed21590e44",
    "plan_code": "BP 2025",
    "code": "RIT123456789",
    "title": "Logistics Cost & Time Management",
    "description": "This project will streamline our internal processes and increase overall efficiency",
    "manager": "John Doe",
    "estimator": "Marie Doe",
    "start_date": "2024-12-01",
    "created_at": "2024-07-29T00:55:50Z",
    "updated_at": null
  },
  "budgets": [
    {
      "budget_id": "024437de-c1ce-4188-97ae-fc0021da157c",
      "portfolio_id": "0821c67e-d331-4c60-9257-75ed21590e44",
      "cost_type": "one_time",
      "description": "External Consulting",
      "comment": "estimativa de consultoria externa",
      "cost_amount": 80000,
      "cost_currency": "EUR",
      "cost_tax": 23.1,
      "amount": 590880,
      "budget_allocations": [
        {
          "year": 2025,
          "month": 3,
          "amount": 221580
        },
        {
          "year": 2025,
          "month": 5,
          "amount": 369300
        }
      ],
      "created_at": "2024-07-29T00:55:50Z",
      "updated_at": null
    },
    {
      "budget_id": "be718119-ae4b-407d-a64a-c7f17acd45f2",
      "portfolio_id": "0821c67e-d331-4c60-9257-75ed21590e44",
      "cost_type": "one_time",
      "description": "Mão de obra do PO",
      "comment": "estimativa do PO",
      "cost_amount": 180000,
      "cost_currency": "BRL",
      "cost_tax": 0,
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
      "created_at": "2024-07-29T00:55:50Z",
      "updated_at": null
    }
  ]
}
```
GET Portfolios FC 02 2025
```json
GET http://localhost:9000/api/v1/portfolios/{{ portfolioIdFC03 }}

HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 29 Jul 2024 15:48:25 GMT
Content-Length: 1427
Connection: close

{
  "portfolio": {
    "portfolio_id": "7670f871-e36a-4174-8888-b2f890442c7c",
    "plan_code": "FC 03 2025",
    "code": "RIT123456789",
    "title": "Logistics Cost & Time Management",
    "description": "This project will streamline our internal processes and increase overall efficiency",
    "manager": "John Doe",
    "estimator": "Marie Doe",
    "start_date": "2025-07-01",
    "created_at": "2024-07-29T00:57:32Z",
    "updated_at": null
  },
  "budgets": [
    {
      "budget_id": "afddfb53-fc28-43aa-9fd5-9592f2b09ef1",
      "portfolio_id": "7670f871-e36a-4174-8888-b2f890442c7c",
      "cost_type": "one_time",
      "description": "External Consulting",
      "comment": "estimativa de consultoria externa",
      "cost_amount": 80000,
      "cost_currency": "EUR",
      "cost_tax": 23.1,
      "amount": 590880,
      "budget_allocations": [
        {
          "year": 2025,
          "month": 10,
          "amount": 221580
        },
        {
          "year": 2025,
          "month": 12,
          "amount": 369300
        }
      ],
      "created_at": "2024-07-29T00:57:32Z",
      "updated_at": null
    },
    {
      "budget_id": "945d8912-1663-4696-8154-2d552e2ee748",
      "portfolio_id": "7670f871-e36a-4174-8888-b2f890442c7c",
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
      "created_at": "2024-07-29T00:57:32Z",
      "updated_at": null
    }
  ]
}
```
