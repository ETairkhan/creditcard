Credit Card Tool

A comprehensive Go-based command-line tool for credit card number validation, generation, and information lookup. This project implements industry-standard algorithms and patterns for working with credit card numbers.
🚀 Features
🔍 Validate

Verify credit card numbers using Luhn's algorithm with support for multiple inputs and stdin:

    Checks number length (minimum 13 digits)

    Validates checksum using Luhn's algorithm

    Supports multiple entries and stdin input

🎲 Generate

Create possible credit card numbers by replacing asterisks with digits:

    Replace up to 4 trailing asterisks

    Generate all possible valid combinations

    Optional random selection with --pick flag

ℹ️ Information

Get detailed information about credit card numbers:

    Card validity status

    Brand identification (Visa, MasterCard, AMEX)

    Issuer lookup

    Configurable brand and issuer databases

🆕 Issue

Generate random valid credit card numbers for specific brands and issuers:

    Customizable brand and issuer requirements

    Ensures valid checksum and formatting

🛠 Technologies Used

    Go (Golang) - Primary programming language

    Luhn's Algorithm - Credit card validation standard

    Command-line Interface - Native Go flag and os packages

    File I/O - For reading brand and issuer configurations

📋 Prerequisites

    Go 1.16 or higher

    Basic understanding of command-line tools

🔧 Installation

    Clone the repository:

bash

git clone <repository-url>
cd creditcard

    Build the project:

bash

go build -o creditcard .

📖 Usage
Validate Credit Card Numbers
bash

# Single validation
./creditcard validate "4400430180300003"

# Multiple validations
./creditcard validate "4400430180300003" "4400430180300011"

# Stdin input
echo "4400430180300003" | ./creditcard validate --stdin

Generate Possible Numbers
bash

# Generate all possibilities
./creditcard generate "440043018030****"

# Pick one randomly
./creditcard generate --pick "440043018030****"

Get Card Information
bash

./creditcard information --brands=brands.txt --issuers=issuers.txt "4400430180300003"

Issue New Card Numbers
bash

./creditcard issue --brands=brands.txt --issuers=issuers.txt --brand=VISA --issuer="Kaspi Gold"

🏗 Architecture

The tool follows a modular command-line architecture:
text

creditcard/
├── main.go              # CLI entry point and command routing
├── validation/          # Luhn algorithm implementation
├── generation/          # Number generation logic
├── information/         # Brand and issuer lookup
└── issue/              # Card issuance logic

Key Components:

    Luhn Validator: Implements the checksum algorithm for credit card validation

    Pattern Generator: Handles asterisk replacement and valid number generation

    Brand Matcher: Matches card numbers against brand prefixes

    Issuer Lookup: Identifies issuing institutions based on card prefixes

📁 Configuration Files
brands.txt Format
text

VISA:4
MASTERCARD:51
MASTERCARD:52
AMEX:34
AMEX:37

issuers.txt Format
text

Kaspi Gold:440043
Forte Black:404243
Halyk Bonus:440563

🎯 Learning Objectives

This project demonstrates:

    Algorithms: Implementation of Luhn's algorithm

    I/O Operations: File reading, stdin/stdout handling

    Data Representation: Efficient number processing and validation

    CLI Development: Professional command-line tool design

🔮 Future Improvements

    Support for additional card brands (Discover, UnionPay)

    Batch processing for large datasets

    JSON output format option

    Rate limiting for generation features

    Extended validation rules (expiry dates, CVV)

⚠️ Known Issues

    Maximum 4 asterisks supported in generation mode

    Brand/issuer files must follow exact format

    No internationalization support
