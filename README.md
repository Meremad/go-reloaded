# Text Processing Tool

## Overview
This project is a command-line tool written in Go that processes text files by applying various transformations. It corrects punctuation, replaces numbers in different numeral systems (hex, bin) with their decimal equivalents, and applies case modifications.

## Features
- **Number System Conversion**: Converts hexadecimal and binary numbers to decimal.
- **Punctuation Fixes**: Corrects spacing around punctuation, removes unnecessary spaces, and fixes quotes.
- **Case Modifications**: Applies uppercase, lowercase, and capitalization transformations based on provided markers.
- **Grammar Fixes**: Corrects the usage of articles "a" and "an" based on the following word.
- **Recursive Processing**: Ensures transformations are applied until the text reaches a stable state.

## Installation
Ensure you have Go installed on your system. Then, clone the repository and build the project:

```sh
# Clone the repository
git clone https://github.com/yourusername/go-text-processor.git
cd go-text-processor

# Build the project
go build -o textprocessor
```

## Usage
Run the program with an input text file and specify an output file:

```sh
./textprocessor input.txt output.txt
```

## Example
### Input:
```
A hour has passed. It is a honor to meet you.
1E (hex) files were added. It has been 10 (bin) years.
" awesome "
```

### Output:
```
An hour has passed. It is an honor to meet you.
30 files were added. It has been 2 years.
"awesome"
```

## File Structure
- **main.go** - Entry point of the application.
- **main_test.go** - Unit tests for the application.
- **proccestext.go** - Core text processing functions.
- **correctness.go** - Functions for punctuation and grammar correction.
- **convert.go** - Functions for numeral system conversions.
- **adjustcousing.go** - Functions for case transformations.

## Testing
To run tests, use:
```sh
go test ./...
```

