package main

import (
    "fmt"
    "time"
)

type samwat struct {
    year, month, day int
}

var BS_YEAR_TO_MONTHS = map[int][13]int{
    2000: {0, 30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
    2001: {0, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
    2002: {0, 31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
    2003: {0, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
    2004: {0, 30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
    2005: {0, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
    2006: {0, 31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
    2007: {0, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
    2008: {0, 31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 29, 31},
    2009: {0, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
    2010: {0, 31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
    2011: {0, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
    2012: {0, 31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30},
    2013: {0, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
    2014: {0, 31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
    2015: {0, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
    2016: {0, 31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30},
    2017: {0, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
    2018: {0, 31, 32, 31, 32, 31, 30, 30, 29, 30, 29, 30, 30},
    2019: {0, 31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
    2020: {0, 31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30},
    2021: {0, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
    2022: {0, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30},
    2023: {0, 31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
    2024: {0, 31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30},
    2025: {0, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
    2026: {0, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
    2027: {0, 30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
    2028: {0, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
    2029: {0, 31, 31, 32, 31, 32, 30, 30, 29, 30, 29, 30, 30},
    2030: {0, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
    2031: {0, 30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
    2032: {0, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
    2033: {0, 31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
    2034: {0, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
    2035: {0, 30, 32, 31, 32, 31, 31, 29, 30, 30, 29, 29, 31},
    2036: {0, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
    2037: {0, 31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
    2038: {0, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
    2039: {0, 31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30},
    2040: {0, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
    2041: {0, 31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
    2042: {0, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
    2043: {0, 31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30},
    2044: {0, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
    2045: {0, 31, 32, 31, 32, 31, 30, 30, 29, 30, 29, 30, 30},
    2046: {0, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
    2047: {0, 31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30},
    2048: {0, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
    2049: {0, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30},
    2050: {0, 31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
    2051: {0, 31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30},
    2052: {0, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
    2053: {0, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30},
    2054: {0, 31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
    2055: {0, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
    2056: {0, 31, 31, 32, 31, 32, 30, 30, 29, 30, 29, 30, 30},
    2057: {0, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
    2058: {0, 30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
    2059: {0, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
    2060: {0, 31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
    2061: {0, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
    2062: {0, 30, 32, 31, 32, 31, 31, 29, 30, 29, 30, 29, 31},
    2063: {0, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
    2064: {0, 31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
    2065: {0, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
    2066: {0, 31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 29, 31},
    2067: {0, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
    2068: {0, 31, 31, 32, 32, 31, 30, 30, 29, 30, 29, 30, 30},
    2069: {0, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
    2070: {0, 31, 31, 31, 32, 31, 31, 29, 30, 30, 29, 30, 30},
    2071: {0, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
    2072: {0, 31, 32, 31, 32, 31, 30, 30, 29, 30, 29, 30, 30},
    2073: {0, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 31},
    2074: {0, 31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30},
    2075: {0, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
    2076: {0, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30},
    2077: {0, 31, 32, 31, 32, 31, 30, 30, 30, 29, 30, 29, 31},
    2078: {0, 31, 31, 31, 32, 31, 31, 30, 29, 30, 29, 30, 30},
    2079: {0, 31, 31, 32, 31, 31, 31, 30, 29, 30, 29, 30, 30},
    2080: {0, 31, 32, 31, 32, 31, 30, 30, 30, 29, 29, 30, 30},
    2081: {0, 31, 31, 32, 32, 31, 30, 30, 30, 29, 30, 30, 30},
    2082: {0, 30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30},
    2083: {0, 31, 31, 32, 31, 31, 30, 30, 30, 29, 30, 30, 30},
    2084: {0, 31, 31, 32, 31, 31, 30, 30, 30, 29, 30, 30, 30},
    2085: {0, 31, 32, 31, 32, 30, 31, 30, 30, 29, 30, 30, 30},
    2086: {0, 30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30},
    2087: {0, 31, 31, 32, 31, 31, 31, 30, 30, 29, 30, 30, 30},
    2088: {0, 30, 31, 32, 32, 30, 31, 30, 30, 29, 30, 30, 30},
    2089: {0, 30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30},
    2090: {0, 30, 32, 31, 32, 31, 30, 30, 30, 29, 30, 30, 30},
}

var ad_scale time.Time = time.Date(1944, time.January, 1, 0, 0, 0, 0, time.UTC)
var bs_scale samwat = samwat {2000, 9, 17}


func convert_ad_to_bs(date_in_ad time.Time) samwat {
    diff := date_in_ad.Sub(ad_scale)
    diff_days := int(diff.Hours()) / 24

    year := bs_scale.year
    day := bs_scale.day + int(diff_days)
    month := bs_scale.month

    for day > BS_YEAR_TO_MONTHS[year][month] {
        day -= BS_YEAR_TO_MONTHS[year][month]

        if month == 12 {
            month = 1
            year += 1
        } else {
            month += 1
        }

    }
    return samwat{year, month, day}
}

func sum(input []int) int {
    sum := 0

    for _, i := range input {
        sum += i
    }
    return sum
}

func convert_bs_to_ad(date_in_bs samwat) time.Time {
    months_of_year := BS_YEAR_TO_MONTHS[date_in_bs.year]
    days := sum(months_of_year[1:date_in_bs.month])
    year := date_in_bs.year - 1

    for year >= bs_scale.year {
        months_for_year := BS_YEAR_TO_MONTHS[year]
        if year == bs_scale.year {
            days += sum(months_for_year[bs_scale.month + 1:])
            days += months_for_year[bs_scale.month] - bs_scale.day
        } else {
            days += sum(months_for_year[1:])
        }
        year -= 1
    }
    hours := days * 24
    return ad_scale.Add(time.Duration(hours) * time.Hour)
}


func main(){
    nepali_date := convert_ad_to_bs(time.Date(2017, time.May, 18, 0, 0, 0, 0, time.UTC))
    fmt.Println(nepali_date)
    fmt.Println(convert_bs_to_ad(samwat{2074, 2, 4}))
}
