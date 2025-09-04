#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include <sys/wait.h>

int main(int argc, char *argv[]) {
    printf("parent: begin\n");
    int rc = fork();
    if (rc < 0) {
        // fork failed
        fprintf(stderr, "fork failed\n");
        exit(1);
    } else if (rc == 0) {
        // child process
        printf("child: running exec\n");
        char *myargs[5];
        myargs[0] = "wc";
        myargs[1] = "-l";
        myargs[2] = "-w";
        myargs[3] = "fwe.c";
        myargs[4] = NULL;
        execvp(myargs[0], myargs);
        printf("child: exec failed\n");
    } else {
        // parent process
        int wc = wait(NULL);
        printf("parent: end\n");
    }
    return 0;
}