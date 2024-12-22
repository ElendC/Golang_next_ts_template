# Quick Refreshing Tutorial
## Some operations
- `=`: Assignment operator
- `==`: Compares values. `1=="1"` and `null==undefined` both returns true.
- `===` Compares value and type. `1=="1"` and `null==undefined` both returns false.

## UseStates
Client side react hook only (Not server).
```ts
const [value, setValue] = useState("Init Value");
   ```
1. First element in the array `value`, is the state. Holds current value and can be named anything
2. Second element `setValue` is a function that updates `value`
3. `useState()` can be initialized (like above) or empty

### Changing value with onClick
```ts
onClick={() => setValue("something")}
```