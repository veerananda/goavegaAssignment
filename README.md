# goavegaAssignment
Assignment 

Create a cron parsing tool

Used Interface and struct to hold methods and data required

The function cronParser(), gives the end output that we can verify even in the test methods

Kind of had a single function that parses whatever message (parseMessage() function) sent to it individually as this saves lot of coding lines

We identify the fields and ranges that are parsed inside this parseMessage function by the range that is sent as input to the function.

Kind of took liberty to send hard coded values for lower and upper ranges of minutes, hours, day of month, month, day of week. Since these are
basic and known things just hardcoded these values and sent as input to the function.

Logic behind the parseMessage function is that we have four symbols "," "/" "*" "-" 
I brokedown the priority of these symbols and splitted the input string based on these symbols.

Most nuclear part of the input expression that we can break is using ",".
Split the input string based on "comma", so that we get all the possible data that user entered.

Now check each element obtained after splitting for / symbol.
If found split to obtain the step values in the range. 
(eg: 1-5/2, after split with respect to "/" symbol , we can traverse the range mentioned 1-5 by skipping two values as denominator is 2)

Now check for "-" symbol and print the range of elements if they fall inside the range for that category.

Finally check for "*" symbol. This symbol doesn't mix with any since it complete range of elements.
This should be single character and the entire range of that category is printed.

Final check is for numbers and words used to describe months or days of week. 
This is handled in the default else part of the parseMessage function.
