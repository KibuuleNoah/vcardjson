# VCard to JSON Converter

## Overview

This project provides a **simple Go utility** to read contacts from a `.vcf` (vCard) file and write them into a **JSON file**.
It focuses on two main operations:

* **ReadVcard** → Extract contacts from a `.vcf` file.
* **WriteContactsToJson** → Save the extracted contacts into a structured `.json` file.

The project defines a lightweight `Contact` struct with the following fields:

* `FullName` – The contact’s full name.
* `Tel` – The contact’s phone number.

---

## Features

* Parses `.vcf` (vCard) files.
* Extracts **full names** and **phone numbers**.
* Cleans phone numbers using regex.
* Outputs contacts in a pretty-printed `.json` file.

---

## Usage

### 1. Clone and Build

```bash
git clone https://github.com/KibuuleNoah/vcardjson.git
cd vcardjson
go build
```


### 2. Run

```bash
go run main.go
```

---

## File Structure

```
project/
│── utils/
│   └── utils.go        # Contains ReadVcard, WriteContactsToJson, ExtractPhoneNumber
│── main.go          # Example entry point using the utils package
```

---

## Example Input (`contacts.vcf`)

```
BEGIN:VCARD
FN:John Doe
TEL;TYPE=cell:+1234567890
END:VCARD
BEGIN:VCARD
FN:Jane Smith
TEL;TYPE=work:+1987654321
END:VCARD
```

## Example Output (`contacts.json`)

```json
[
 {
  "fullname": "John Doe",
  "tel": "+1234567890"
 },
 {
  "fullname": "Jane Smith",
  "tel": "+1987654321"
 }
]
```

---

## Requirements

* Go 1.18+
* A valid `.vcf` file with contacts.

---

## Notes

* If the JSON output file already exists, it will be **overwritten**.
* Regex ensures only digits (and symbols after `:`) are extracted from phone fields.
* Designed for **simplicity and clarity** – a good starting point for larger contact management tools.
