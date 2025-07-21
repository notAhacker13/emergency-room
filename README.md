# emergency-room
Emergency room implementation in a Hospital using priority queue. Implemented in Go with unit tests.


## Usage

### 1. Clone the repository

```sh
git clone https://github.com/notAhacker13/emergency-room.git
cd emergency-room
```

### 2. Run Tests

```sh
go test
```

All tests should pass, verifying the correctness of the priority queue logic.

## Example

Here’s a quick example of how the Emergency Room works:

```go
er := NewEmergencyRoom()

er.AdmitPatient(&Patient{name: "John", severity: Moderate})
er.AdmitPatient(&Patient{name: "Sam", severity: Critical})

next := er.HandleNextPatient()
// next will be "Sam" because Critical > Moderate
```

## How It Works

- **AdmitPatient**: Adds a new patient to the queue.
- **HandleNextPatient**: Pops the patient with the highest severity.
- **UpdatePatientStatus**: Changes a patient’s severity and reorders the queue.

## Continuous Integration

This project uses [GitHub Actions](.github/workflows/go.yml) to automatically run tests on every push and pull request.

## License

MIT (or specify your preferred license)

---

*Made with Go and ❤️ by [notAhacker13](https://github.com/notAhacker13)*
