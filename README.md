# zero-balance-budget-api
A monolithic API to get a working proof-of-concept before tackling problems like container orchestration and distributed permissions.

## Purpose

### Zero-Balance Budget
The idea here is to configure a budget/plan, and then take each paycheck and automatically calculate what goes toward expenses, bills, and savings. Eventually, it may let you upload and categorize purchases to see which areas need more or less budget.

### Why a Monolith? It's 2018.
I started writing my own microservices, but it turns out tackling JWT-based authentication, permissioning/authorization, inter-service communication, and worrying about deployment all at once is a LOT to learn/plan at once. This lets me get to an iteration zero and start budgeting before tackling those less-necessary problems.
