// Define a global variable 'Module' with a method 'onRuntimeInitialized':
Module = {
  onRuntimeInitialized() {
    // this is our application:
    console.log(cv.getBuildInformation());
  },
};
// Load 'opencv.js' assigning the value to the global variable 'cv'
cv = require('./opencv.js');
