<p align="center">
<img src="internal/assets/logo_medium.png" alt="logo" width="110" height="110">
</p>
<h1 align="center"><a href="https://pkg.go.dev/github.com/yinebebt/ethiopiancalendar">Ethiopian Calendar (ባሕረ-ሐሳብ)</a></h1>

![build-workflow](https://github.com/yinebebt/ethiopiancalendar/actions/workflows/build-and-test.yml/badge.svg)

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

* <p><a href="https://pkg.go.dev/github.com/yinebebt/ethiopiancalendar"  target="_blank" >pkg.go.dev</a></p>

## Installation
Install using below go command:
```bash
go get github.com/yinebebt/thiopiancalendar
```

## Author
- [Yinebeb T.](https://github.com/yinebebt)
