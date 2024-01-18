# Drought

Due to the continuous drought that happened recently in some regions of Brazil, the Federal Government created an agency to assess the consumption of these regions in order to verify the behavior of the population at the time of rationing. This agency will take some cities (for sampling) and it will verify the consumption of each of the townspeople and the average consumption per inhabitant of each town.
### Input
The input contains a number of test's cases. The first line of each case of test contains an integer N (1 ≤ N ≤ 1 * 10 6), indicating the amount of properties. The following N lines contains a pair of values X (1 ≤ X ≤ 10) and Y ( 1 ≤ Y ≤ 200) indicating the number of residents of each property and its total consumption (m3). Surely, no residence consumes more than 200 m3 per month. The input's end is represented by zero.
### Output
For each case of test you must present the message “Cidade# n:”, where n is the number of the city in the sequence (1, 2, 3, ...), and then you must list in ascending order of consumption, the people's amount followed by a hyphen and the consumption of these people, rounding the value down. In the third line of output you should present the consumption per person in that town, with two decimal places without rounding, considering the total real consumption. Print a blank line between two consecutives test's cases. There is no blank line at the end of output.



| Input Sample | Output Sample  |
| ------ | ------ |
| 3 <br> 3 22 <br> 2 11 <br> 3 39 <br> 5 <br> 1 25 <br> 2 20 <br> 3 31 <br> 2 40 <br> 6 70 <br> 2 <br> 1 1 <br> 3 2 <br> 0 |Cidade# 1: <br> 2-5 3-7 3-13 <br> Consumo medio: 9.00 m3. <br><br> Cidade# 2: <br> 5-10 6-11 2-20  1-5 1-25 <br> Consumo medio: 9.00 m3. <br><br>  Cidade# 2: <br> 3-0 1-1 <br> Consumo medio: 0.75 m3.|

