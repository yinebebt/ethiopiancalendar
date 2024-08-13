<p align="center">
<img src="internal/assets/logo_medium.png" alt="logo" width="110" height="110">
</p>
<h1 align="center"><a href="https://pkg.go.dev/github.com/Yinebeb-01/ethiopiancalendar">Ethiopian Calendar (ባሕረ-ሐሳብ)</a></h1>

![build-workflow](https://github.com/Yinebeb-01/ethiopiancalendar/actions/workflows/build-and-test.yml/badge.svg)

## Description
The Ethiopian calendar(ባሕረ-ሀሳብ) is used to get Fasting and Holiday's specific date with in a year based on 
[EOTC](https://www.ethiopianorthodox.org/) calendar. It also designed to facilitate the conversion between Ethiopian dates (in the format yy-mm-dd) and 
Gregorian dates. Ethiopia follows its own calendar system, which consists of 13 months, each with 30 days. 

### Functionality
This API allows users to:
* Fetching all festivals with a year such as year, Basic data and Fasting dates.
* Convert Ethiopian dates to Gregorian dates.
* Convert Gregorian dates to Ethiopian dates.

### Usage
To utilize the API, simply send a date or year using the specified endpoints. For the date conversion,the API will 
respond with the converted date in JSON format whereas for the calendar, it will respond data in a json object.

#### Documentation

* <p><a href="https://ethiocal.koyeb.app/v1"  target="_blank" >Overview</a></p>

* <p><a href="https://ethiocal.koyeb.app/v1/swagger/index.html"  target="_blank" >Try on swagger</a></p>

* <p><a href="https://pkg.go.dev/gihub.com/ethiopiancalendar"  target="_blank" >pkg.go.dev</a></p>


v1 Convert Gregorian DOverview Ethiopian Date

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

3. Get Ethiopian Calendar such as Holidays and Fasting dates.

```curl
GET /bahire-hasab/{year}
```

* Parameters:

  `year:` The Ethiopian year.

**Example**:

```curl
GET /bahire-hasab/2016
```

Response:

```json
{
  "data": {
      "year": {
        "year":2016,
        "evangelist_num":0,
        "evangelist":"ዮሐንስ(John)",
        "day_of_the_week":"ማክሰኞ(Tuesday)"
      },
      "basic": {
        "medeb":11,
        "wenber":10,
        "abektie":20,
        "metiq":10,
        "beale_metiq":2,
        "mebaja_hamer":18,
        "nenewie":{
          "date_of_the_month":18,
          "month_of_the_year":6
        }
      },
    "fasting":{
      "abiy":{
        "date_of_the_month":2,
        "month_of_the_year":7
      },
      "debrezeit":{
        "date_of_the_month":29,
        "month_of_the_year":7
      },
      "hosanna":{
        "date_of_the_month":20,
        "month_of_the_year":8
      },
      "siklet":{
        "date_of_the_month":25,
        "month_of_the_year":8
      },
      "tinsaye":{
        "date_of_the_month":27,
        "month_of_the_year":8
      },
      "rkbekahnat":{
        "date_of_the_month":21,
        "month_of_the_year":9
      },
      "dihnet":{
        "date_of_the_month":19,
        "MonthOfTheYear":10
      },
      "hawariyat":{
        "date_of_the_month":17,
        "month_of_the_year":10
      },
      "erget":{
        "date_of_the_month":6,
        "month_of_the_year":10},
      "peraklitos":{
        "date_of_the_month":16,
        "month_of_the_year":10
      }
    }
  }
}
```

## Installation
Install using below go command:
```bash
go get github.com/yinebebt/thiopiancalendar
```

## Author
- [Yinebeb T.](https://github.com/yinebebt)
