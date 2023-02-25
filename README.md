# Money App

The Money App is a App for budgeting, analysing and visualization of wealth, income and spending.
It is based around transaction and stores all it's information in form of generic transactions, which can be imported from diffenrent banks.
The places where wealth can be stored is called a pod.
A transaction is always one of three different types (IN, OUT, MOVE) depending wheater the receiver, the sender or both are pods of the user.
Each transaction can have multple budgets, inbudgets and debts and arbitrary other tags.

The App is implemented through a Golang-Backend and a Svelte-Kit-Frontend.
For them to work together some sort of authentication proxy (like [proxauth](https://github.com/xilefmusics/proxauth)) is needed since the backend doesn't handle user authentication itself.

## Backend

The backend stores the list of transactions per user.
It selects the data it operates on through a header `user: <username>`.
If a request changes data it adds a change event to a list of events.
The newest event can be rolledbacked through the endpoint `/undo`.

### GET /transactions

Returns a list of transactions.

| Filter   | Default |
| -------- | ------- |
| year     | \*      |
| month    | \*      |
| pod      | \*      |
| debt     | \*      |
| budget   | \*      |
| inbudget | \*      |
| type     | \*      |
| id       | \*      |

### POST `/transactions`

### `PUT /transactions`

### DELETE /transactions

| Endpoint           | Method | Description                                                      |
| ------------------ | ------ | ---------------------------------------------------------------- |
| /transactions      | GET    | Returns all transactions.                                        |
| /transactions      | POST   | Updates a transaction based on IDs.                              |
| /transactions      | PUT    | Creates new transactions, ignores if they are not yet there.     |
| /transactions      | DELETE | Deletes the the given list of transactions based on their IDs.   |
| /lint              | GET    | Returns all linting problems of transactions.                    |
| /reindex           | GET    | Sorts transactions by date and overwrites the ID with the index. |
| /undo              | GET    | Undoos the last action. Returns the applied event.               |
| /pods              | GET    | Returns all pods.                                                |
| /debts             | GET    | Returns all debts.                                               |
| /budgets           | GET    | Returns all budgets.                                             |
| /inbudgets         | GET    | Returns all inbudgets.                                           |
| /tags              | GET    | Returns all tags.                                                |
| /history/debt      | GET    | Returns the history of debts.                                    |
| /history/budgets   | GET    | Returns the history of budgets.                                  |
| /history/inbudgets | GET    | Returns the history of inbudgets.                                |
| /history/pod       | GET    | Returns the history of pods.                                     |
| /history/wealth    | GET    | Returns the history of wealth.                                   |
