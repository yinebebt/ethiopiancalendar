<p align="center">
<img src="internal/assets/logo_medium.png" alt="logo" width="110" height="110">
</p>
<h1 align="center"><a href="https://pkg.go.dev/gitlab.com/Yinebeb-01/ethiopiandateconverter">Ethiopian Date Converter</a></h1>

## Description
The Ethiopian Date Converter API is designed to facilitate the conversion between Ethiopian dates
(in the format yy-mm-dd) and Gregorian dates. Ethiopia follows its own calendar system, which consists of 13 months,
each with 30 days. 

### Functionality
This API allows users to:
* Convert Ethiopian dates to Gregorian dates.
* Convert Gregorian dates to Ethiopian dates.

### Usage
To utilize the API, simply send a date using the specified endpoints. The API will respond with the converted date 
in JSON format.

#### Example Usage

1. Convert Gregorian Date to Ethiopian Date

```curl
GET /ad-to-et/{date}
```

* Parameters:

    `date:` The Gregorian date to convert (in _yy-mm-dd_ format).

**Example**:

```curl
GET /ad-to-et/2013-09-11
```

Response:

```json
{
"gregorian_date": "2021-05-22"
}
```

2. Convert Gregorian Date to Ethiopian Date

```curl
GET /et-to-ad/{date}
```

* Parameters:

    `date:` The Ethiopia date to convert (in _yy-mm-dd_ format).

**Example**:
```curl
GET /et-to-ad/2021-05-22
```
Response:
```json
{
"ethiopian_date": "2013-09-11"
}
```

## Installation

Install using below go command:
```bash
go get gitlab.com/Yinebeb-01/ethiopiandateconverter
```

## Author
- [Yinebeb T.](https://github.com/Yinebeb-01/)
- [python version](https://github.com/dimagi/ethiopian-date-converter)