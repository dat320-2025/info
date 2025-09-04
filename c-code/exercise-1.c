// write a c program that prints every 5 seconds and stops after 30 seconds
// the program should print the current time in the format HH:MM:SS and the elapsed time in seconds and process id


#include <stdio.h>
#include <time.h>
#include <unistd.h>

int main() {
    time_t start_time = time(NULL);
    time_t current_time;
    int elapsed_time;
    pid_t process_id = getpid();

    while (1) {
        current_time = time(NULL);
        elapsed_time = (int)(current_time - start_time);

        // Print current time in HH:MM:SS format
        struct tm* timeinfo = localtime(&current_time);
        printf("%02d:%02d:%02d ", timeinfo->tm_hour, timeinfo->tm_min, timeinfo->tm_sec);

        // Print elapsed time in seconds and process id
        printf("Elapsed Time: %d seconds, Process ID: %d\n", elapsed_time, process_id);

        // Sleep for 5 seconds
        //use spinlock to sleep for 5 seconds
        
        sleep(5);

        // Check if 30 seconds have passed
        if (elapsed_time >= 30) {
            break;
        }
    }

    return 0;
}