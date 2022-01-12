[![License](https://img.shields.io/badge/License-Apache_2.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

<p align="center">
    <br>
    <img src="https://raw.githubusercontent.com/cohere-ai/tokenizer/readme-initial-draft/cokenizer.png" width="600"/>
    <br>
<p>

# About

Cohere's `tokenizers` library provides an interface to encode and decode text given a computed vocabulary, and includes pre-computed tokenizers that are used to train Cohere's models. 

We plan on eventually also open sourcing tools to create new tokenizers. 

## Installation 
...

## Example using Go
Choose or create a tokenizer inside of the tokenizers folder including both a encoder.json file and a vocab.bpe file and instantiate an encoder as seen below:
```
import (
  ...
  "https://github.com/cohere-ai/tokenizer"
)

encoder := NewFromPrebuilt("coheretext-50k")
```
From here, input the string as a parameter to the encoder's NumTokens(string) method to find how many tokens the string is comprised of in the form of an integer:
```
fmt.Printf("%d", encoder.NumTokens("Example String"))
# Output: 2
```
Or find out exactly what tokens are being used using the ListTokens(string) method, which returns a slice of strings:
```
fmt.Printf("%v", encoder.ListTokens("Example String For Listing Tokens")
# Output: ["Example" " String" " For" " Listing" " Tok" "ens"]
```

## Speed
Using a 2.5GHz CPU, encoding 1000 tokens takes approximately 6.5 milliseconds, and decoding 1000 tokens takes approximately 0.2 milliseconds.
