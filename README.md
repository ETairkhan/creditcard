# Credit Card Utility CLI

A versatile command-line interface (CLI) tool for generating, validating, and obtaining information about credit cards. This application is designed for developers, testers, and businesses needing reliable utilities for handling credit card data.

## Features

- **Validate**: Check if credit card numbers are valid using the Luhn algorithm.
- **Generate**: Generate random valid credit card numbers for testing purposes.
- **Information**: Retrieve brand and issuer details for a given credit card number.
- **Issue**: Generate credit card details for a specific brand and issuer.

---

## Installation

1. Clone this repository:

```bash
git clone https://github.com/your-username/creditcard-cli.git
```
2. Navigate to the project directory:

```bash
git clone https://github.com/your-username/creditcard-cli.git
```

3. Build the binary:
```bash
go build -o creditcard
```

## Usage

**1. Validate**

Check if a credit card number is valid.

- From stdin:
```bash
echo "4400430180300003" "4400430180300011" | ./creditcard validate --stdin
```
- Direct input:
```bash
./creditcard validate 4400430180300003 4400430180300011
```

**2. Generate**

Generate random valid credit card numbers.

- Using a pattern:

```bash
./creditcard generate <pattern>
```

Example: Generate a card number starting with 4400:

```bash
./creditcard generate 4400
```
- Random valid number:
```bash
./creditcard generate --pick
```

3. Information

Retrieve card brand and issuer details.

From stdin:
```bash
echo "4400430180300003" | ./creditcard information --stdin --brands brands.txt --issuers issuers.txt
```
Direct input:
```bash
./creditcard information --brands brands.txt --issuers issuers.txt 4400430180300003
```

4. Issue

Generate card details for a specific brand and issuer.
```bash
./creditcard issue --brands brands.txt --issuers issuers.txt --brand Visa --issuer
```

## Flags

- `--stdin`: Read input from standard input.
- `--brands`: Path to `brands.txt` file (used for `information` and `issue` features).
- `--issuers`: Path to `issuers.txt` file (used for `information` and `issue` features).
- `--pick`: Generate a random valid card number.
- `<pattern>`: Generate a card number matching a specific pattern.


## File Structure
```bash
.
├── main.go            # Entry point of the CLI application
├── validate.go        # Functions for validating credit card numbers
├── generate.go        # Functions for generating card numbers
├── information.go     # Functions for retrieving brand and issuer info
├── issue.go           # Functions for issuing credit card details
├── helpers.go         # Utility functions
├── brands.txt         # List of card brands
├── issuers.txt        # List of card issuers
└── README.md          # Documentation
```

## Error Handling

- All errors are logged and handled gracefully using `defer` and `recover`.
- Invalid inputs will display helpful error messages to guide the user.

---

## Contributing

Contributions are welcome! Feel free to submit issues or pull requests to improve the project.

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
