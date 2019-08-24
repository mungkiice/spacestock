# Data Manipulation

## 1. Working with Arrays
### 1.1. Simple with sort
```javascript
husband = ["alphard", "civic", "terios"]
wife = ["terios", "jazz"]

//Write a merge function
function merge(arr1, arr2) {
    //expected: ["alphard", "civic", "jazz", "terios"]
}
```

### 1.2. Object merge
```javascript
invoices = [
    {"to": "John Doe", "amount": 120000},
    {"to": "John Doe", "amount": 80000},
    {"to": "Jane Doe", "amount": 100000}
]

//write a function to group invoices by the person
function groupInvoiceByPerson(invoices) {
    /** expected: [
        {"to": "John Doe", "amount": 200000},
        {"to": "Jane Doe", "amount": 100000}
    ]**/
}
```

### 1.3. Finding similar string
```javascript
bannedWords = ["scam", "spam", "dirty"]
input1 = "this is a scammer"
input2 = "this is a spam message"
input3 = "this contains dirty word"
input4 = "this is a valid message"

//write a function which returns true if a string contains banned words
function isBanned(words, bannedWords) {
    //input1 returns true
    //input2 returns true
    //input3 returns true
    //input4 return false
}
```

# Golang Specifics

## 1.1. Struct
What does the john.Name prints and why?

```
type Human struct {
    Name string
}

func ChangeName(human Human, to string) {
    human.name = to
}

func main() {
    john := Human{Name: "John Doe"}
    ChangeName(human, "Jane Doe")

    //What does the john.Name prints and why?
    fmt.Println("%s", john.Name)
}
```

## 1.2. Please setup a REST service for CRUD

endpoints: 
1. [GET] /apartment
2. [POST] /apartment
3. [PUT] /apartment/{id}
4. [DELETE] /apartment/{id}

For database, you can either
1. Mock the database
2. Use docker

Or whatever database you are familiar with
But we prefer to have mongo / postgre

And add some unit test to the service