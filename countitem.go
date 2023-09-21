package main

func countItem(inventory []string, item string) int {
    count := 0
    for _, i := range inventory {
        if i == item {
            count++
        }
    }
    return count
}