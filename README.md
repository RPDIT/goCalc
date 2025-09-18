# Terminal Calculator

A CLI calculator that was utilized as a tool to explore deeper aspects of Go, its errors, and its testing.

I started this project by building the basic functions of the calculator and considering how I wanted to structure the application. By determining I wanted each function to return a Result structure that has both an error and a value, I was able to create uniformity across the implementation of the code in my application. This same uniformity then allowed me to easily explore the deep possilities of testing in Go. While I did not have much experience with Go's tests previously, taking this opportunity to write assertions and handle the results gave me a crash course on the flexability and confidence that firmly tested software provides.

# Challenges

When the tests are complete and able to run, they are immensely powerful; but, with my exploration of Go's testing suite, I also came to understand why simple tests can be the most beneficial. Writing out multiple tests with complex value and assertion structure can be enticing, but simple, hard-coded, tests may be the most flexible and sustainable as these projects scale.

While this is a learning endeavor for me, I make a personal note to not just seek out code blocks from other sources and paste them. I will commonly search up concepts or questions I run into, but I do not wish to implement code that I cannot understand. With this preface, I ran into a minor roadblock coming up with a solution to increasing a number to the nth power, or exponential increase. While the math is easy to understand for a human, it took me some thinking to understand how to multiply x by itself n times. I settled on iteration through a range, increasing a counter value and muliplying it by the root each time. This solution works but I feel there is a faster option, and reviewing Go's math package documentation shows that there really are levels to this idea. There are many edge cases I do not account for and this is a bit of a ham-fisted solution, lacking handling of the nuance of negative and 0 squares.

# Closing Thoughts

I have gained a great deal through this project and will continue to increase my understanding of both coding and this language at each step. If this is a project I plan to return to, there are a few aspects I would address: edge case handling, add a nth root function, process and more effectively wrap errors, and consider adding a GIO GUI.
