# **U**ser In**put**

**U**ser in**put**,or **uput** (probably best to avoid ui, since it would
collide with user interface) is a library built to handle input in a
streamlined way from command-line.

Aim at being minimal and flexible enough to provide user input
functionality from basic scripts to monolithic applications. The library
aims at being as feature rich and unimpiniated as possible to acheive
maximum flexibility. The library strives to only be opiniated on matters of
security.

**Features include:**

* Modular functionality by relying heavily on sub-packages that can be
  imported individually in small projects without creating a large footprint.
* User prompts for a variety of data types including: boolean, string list, string multiselect, string path, and more...
* Localization (i18n) support that is not opininated and able to play nicely with existing implementation.
* Pre-validation and Post-validation transformations to prepare data for validation and saving into databases.
* Support for *single variables* and *structs* using supporting two different methods, built from researching existing popular implementations.
* Variety of specific type validation within the builtin data types; ranging from common to niche, impelemented as sub-packages to be imported as needed.
* Relying on stdlibs like unicode over heavy reliance on regex.
* Focus on security, translating to validation of all printed variables
  within the library (for example Error messages). Great care has been
taken to research for best ways to validate types of data (for example, to avoid look-a-likes when comparing UTF8 strings).

All functionality will be implemented as sub-packages, with only a minimal
set of functionality required to use any specific sub-package. For
example, if a script only has string input, the developer can include
only string validations and string prompt sub-packages, which will call
the miminimal set of common code and be ready to use with very little
footprint added to the project.

The library is designed around the fact that messages would be localized
but also delegates the responsibilty of how localizaiton *(i18n)* will occur to
other parts of the application to acheive higher flexibility. No i18n
libraries are used, the developer passes their messages.

Proper user input procedure goes beyond just a pretty prompt, it
requires a consistent pipeline that the developer can hook functions
into.

*(User Prompt)* -> *(Pre-Validation Transformation)* -> *(Input Data Validation)* -> *(Post-Validation Transformation)*

**Pre-Validation Transformation** allows the developer to prepare data for
comparison, for example, one may choose to downcase all emails before
making comparisons to avoid storing duplicate records with varrying
lettercases.

**Post-Validation Transformation** allows the developer to prepare data
for saving or use before returning it from the user input package. Common
transformations can be loaded as sub-packages from the library or custom
transformations can be loaded.


## Command-line prompts

A collection of consistent command-line prompts are included as a
sub-package that can be included, so this library can be used for a wide
variety of user input applications.

Command line prompts are differentiated based on the type of data that
is being input by the user, for example: a boolean type would require a
Yes/No prompt.

When defining a prompt, one must provide:

1. The data type of the user input.
2. The default value.
3. If the value is required.
4. How to handle a failure to validate.


### Data Type Prompts

**Boolean** provides a (Y)es/(N)o prompt

**String (List)**

**String (List, with multiple selectable options)**

**String (Path)**


## **Valid**ation
Validation is technically a separate but closely tied subject because
any user input must be validated, even if the user is trusted. This
design philosophy is known as 'security-in-depth' and assumes all
trusted parties have been compromised and all security precuations have
failed. This entire library and all associated projects under the
Hackwave organization are built with the same security first approach.

It is technically separate because even though every user input must be
validated, validation is also important for non-user input like sensor
data, or calculated data.

Many validation libraries exist, and several input prompt libraries
exist but they are all missing one thing are another, so after reviewing
all existing available projects on both Github and Gitlab I have
collected the best features, implemented them in a modular way so that a
minimal set of features could be used without including the entire set.

Of the existing libraries, the most obvious way to differentiate
impelementations is if validation is aimed at structs or individual variables.

**uput** aims to provide both, in a way that they can be included
separately, to fit the needs of the developer and avoid codebase bloat.

**Input, Validation, Transform Issues And Security**
Below are a few common issues met by developers of similar libraries or common security concerns within the topics of user input, validation, and transforms. Thoughtful research in the subject can avoid libraries re-implementing mistakes already experienced by previous related projects:

[Validation functions return type error, to avoid err is always != nil.](https://stackoverflow.com/questions/29138591/hiding-nil-values-understanding-why-golang-fails-here/29138676#29138676)

[Transforms/Normalization as part of the validation pipeline to prevent
security issues associated with look-a-likes. Example: The character K,
and the character K for Kelvin
Temperature](https://blog.golang.org/normalization)

Each is a sub-package (user input, validation, and transformation), allowing it to be used exclusively, or together.


### Reviewing Existing Implementations
After reviewing existing validation libraries available for Go, a few common strategies emerge as the most popular methodologies for validating data in Go.

In general, struct based validation is the most popular method of
validation. Within that category, there are two major strategies for
defining validation schemas for struct based validation.

### Struct Validation
**Struct: Struct Tag Based Schema Declaration**
These methods use a **Tag or Separate Object (or Rules) Based Schema Validation** to define an entire object (struct) to be validated. The primary difference between the two is the way the validation schema is being declared.

The most popular method to declare a validation schema for a struct is
using *existing struct tag system* to declare validations inline with
the Field declaration when defining a struct. One may be familiar with
tags, they are most often used when working with marshalling "json", "yaml" or "toml" tags when loading or saving configuration files. An Example of both "json" tag and our custom "validate":

      type User struct {
        Id       int     `json:id,validate:length:10`
        Username string
`json:username,validate:between:2-24;alphanumeric`
      }

However, there are legitimate criticisms of this strategy, and while it
may be perfect for some projects, it is not ideal for all.

**Struct: Object (or Rules Object) Based Schema Declaration**
Alternatively, the next most popular method to define a validation schema is using *a "validation"-er object* to define a set of rules. Passing a struct to a "validate"-er will typically return a slice of all validation errors.

Ideally these type of impelemtnations would allow for schemas to be exported generating a specification for the developed software's API. Yet, no library seems to actually implement an ability to export a complete JSON or YAML document that includes all declared object schema to automate documentation generation for APIs.

Each methodology has benefits and because of this, to avoid being too
opinionated, each methodology is implemented as sub-packages, allowing the developer to choose one or both methods.


### Single Variable Validation
And while these methods work well for objects, for single variables, they feel too heavy and clunky. So an additional method has been implemented to cater to developers who need validation for individual variables and not just validation for objects.

Inline declaration of validation schemas that support inline passing of single variables of any type to be validated, empowers developers to validate any variable, even in tiny scritps, utilizing a minimal one line validation.

**Chainable Inline Schema Definition**
By providing validation with only a single human readable line of code that immediately returns a validated/transformed(if transformed) input string and slice of validation errors (if any). Enabling the developer to do a quick:

    // API is not solidified, and subject to change
    isValid, userInput, errs := valid.IfString("test string").IsContaining("cool").IsLessThan(5).IsIn([]string{"test", "best", "mega"}).NoNumeric().IsEmpty().IsValid()
    if isValid {
	    valid.PrintErrors(errs)
    } else {
	    fmt.Println("validated user input is:", userInput)
    }

Intentionally unopiniated, utilizes no special library defined data types, only normal builtin Go datatypes are returned, and this is acheived without feature loss: it still supports localization (i18n) of error messages.

This method works well for one-off validation of a variable inline in a simple script for example.

This method enables developers to easily define their own custom validations or even custom input, output and transform functions.

**Validation Functions** for a given data type, for example, *string*:

    [ IsEmpty(), IsBetween(4, 255), IsNumeric(), NotAlphabetic(), ... ]

Are called in a chain between the **Input Functions**:

    [ IfString(s string), IfInt(i int), IfUint(u uint), IfMap(m map[string]string) ]

And finalizing with an **Output function** which returns *bool*,
dataValue, and any and all errors so they can all be printed to clarify
what exactly needs to be changed to acheive successful validation:

    [ IsValid(errorMessages map[string]string) (bool, interface{}, []error) ]

After the input function and before validation, **transformation functions**
can be called to normalize data, preparing it for validation, or after
validation to prepare data for use, like saving into a DB.

Beyond starting with input and ending with output, there are no strict
rules and any number of validations and transformations can be called.

*The API is not solidified yet, as the package is still under active
development, and is subject to change. If you have suggestions, please
create an issue to initiate a discussion on relevant topics.*

## Transformations
Transformations are necessarily to prepare data for validation or
prepare data for use after validation.

For example, [secure username implementation in a UTF8 environment requires normalization of characters to pretent look-a-likes](https://blog.golang.org/normalization).

## Under Development
Not all features are implemented yet, as of writing only the string
validations have begun active development. The API is not stabilized,
the library is not yet 0.1.0 and is subject to change. Pull requests are
welcome.

