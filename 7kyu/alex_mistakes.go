// Alex is transitioning from website design to coding and wants to sharpen his skills with CodeWars.
// He can do ten kata in an hour, but when he makes a mistake, he must do pushups. These pushups really tire poor Alex out, so every time he does them they take twice as long. His first set of redemption pushups takes 5 minutes. Create a function, alexMistakes, that takes two arguments: the number of kata he needs to complete, and the time in minutes he has to complete them. Your function should return how many mistakes Alex can afford to make.

package main

func AlexMistakes(numberOfKatas, timeLimit int) int {
  // 10 kata per hour = 6 minutes per kata
  timeForKatas := numberOfKatas * 6
  remainingTime := timeLimit - timeForKatas
  
  mistakes := 0
  pushupDuration := 5
  
  // Count how many mistake sets fit in remaining time
  for remainingTime >= pushupDuration {
    remainingTime -= pushupDuration
    mistakes++
    pushupDuration *= 2 // Each set takes twice as long
  }
  
  return mistakes
}
