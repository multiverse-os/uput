# **U**ser In**put**

**U**ser in**put**,or **uput** (probably best to avoid ui, since it would
collide with user interface) is a library built to handle input in a
streamlined way from command-line.

Aim at being minimal and flexible enough to provide user input
functionality from basic scripts to monolithic web applications

**Features include:**

* Modular functionality by relying heavily on sub-packages that can be
  imported individually in small projects without creating a large footprint.
* User prompts for a variety of data types including: boolean, string list, string multiselect, string path, and more...
* Localization (i18n) support that is not opininated and able to play nicely with existing implementation.
* Pre-validation and Post-validation transformations to prepare data for validation and saving into databases.
* Support for *single variables* and *structs* using supporting different methods
* Variety of type validation; ranging from common to niche, impelemented as sub-packages to be imported as needed.

All functionality will be implemented as sub-packages, with only a minimal
set of functionality required to use any specific sub-package. For
example, if a script only has string input, the developer can include
only string validations and string prompt sub-packages, which will call
the miminimal set of common code and be ready to use with very little
footprint added to the project.

The library is designed around the fact that messages would be localized
but also delegates the responsibilty of how localizaiton *(i18n)* will occur to
other parts of the application to acheive higher flexibility.

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

Of the existing libraries there the first main distinction is if
validation is for structs or for individual variables.

**Common issues**

[Validation functions return type error, to avoid err is always != nil.](https://stackoverflow.com/questions/29138591/hiding-nil-values-understanding-why-golang-fails-here/29138676#29138676)


Each is a sub-package, allowing it to be used exclusively, requiring a
minimum amount of additional code.


### Common Implementations
After reviewing existing code there are a few different strategies to impelement validation functionality. 

**Schema**
The first method I will call schema or rule based validation. This method is implemented so a developer using it creates a "validation"-er object, then uses this set or rules or schema against individual values or objects (structs). 

The two main strategies for defining a schema is by:

1. Tags defined inside a struct declaration
2. "validation"-er object declaration inline

Each of these methods of defining have their own positives and negatives.

In general, methodology can be very clean, human readable and is the most popular. It allows one to save a definition with the object and repeatedly use the same schema over and over. Ideally these type of impelemtnations would allow for schemas to be exported generating a specification for the developed software's API (I have not seen this actually impelmented yet). 

I believe there are benefits to this strategy when working on objects, but for individual variables, it feels clunky.

**Chaining**
Chaining feels like the best method for one-off validation of a variable inline in a simple script. It can keep a minimal footprint by only importing necessary validations being used and it keeps the logic being added very simple and human readable. 

There are limitations but one could fallback to schema based validation if the software is complex enough to warrant it.


### Single Variable Validation
The best way to handle single variable validation is by having a
recursively returned validation function that allows chaining.


      isValid, userInput, errs := valid.IfString("t0").IsContaining("cool").IsLessThan(5).IsIn([]string{"test", "best", "mega"}).NoNumeric().IsEmpty().IsValid()
			if len(errs) > 0 {
				valid.PrintErrors(errs)
			} else {
				fmt.Println("validated and usable userInput is:", userInput)

### Struct Validation
There are two common methods to initializing and specifying struct based
validations, the most common method is **Tag based**, relying on the
existing tag system within the struct, which is nice because it is
consistent with other functionality.

A few other libraries like ozzo prefers a **Schema based** *(sometimes
called rules)* which is implemented by defining a schema separate from
the struct.

Each methodology has benefits and because of this, each methodology is
implemented as sub-packages, allowing the developer to choose one or both
methods.


## Transformations
Transformations are necessarily to prepare data for validation or
prepare data for use after validation.





