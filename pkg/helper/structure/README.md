# Structure Package

The Structure package provides a set of helper functions and interfaces to facilitate data conversion, manipulation, and structure management within the application.

## Overview

This package offers utilities for:
- **Casting:** Converting data types and structures.
- **Conversion:** Defining generic conversion interfaces with both standard and MongoDB-specific implementations.
- **Manipulation:** Providing helper functions and implementations to manipulate data structures effectively.

## Files

- **[caster.go](pkg/helper/structure/caster.go):** Functions and interfaces for data casting operations.
- **[converter.go](pkg/helper/structure/converter.go):** Defines the converter interface for data conversion.
- **[converter_std_impl.go](pkg/helper/structure/converter_std_impl.go):** Standard implementation of the converter interface.
- **[function_converter.go](pkg/helper/structure/function_converter.go):** Provides a functional approach to data conversion.
- **[function_converter_mongo_impl.go](pkg/helper/structure/function_converter_mongo_impl.go):** MongoDB-specific converter implementation using functional converters.
- **[manipulater.go](pkg/helper/structure/manipulater.go):** Interfaces and base implementations for data manipulation.
- **[manipulator_std_impl.go](pkg/helper/structure/manipulator_std_impl.go):** Standard implementation of the manipulation interface.
- **[manipulater_helper.go](pkg/helper/structure/manipulater_helper.go):** Helper functions to assist with data manipulation tasks.

## Usage

Import the package into your Go project:

````go
import "path/to/pkg/helper/structure"