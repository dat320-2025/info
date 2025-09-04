#include <stdio.h>
#include <sys/wait.h>
#include <unistd.h>
#include <stdlib.h> // Include the <stdlib.h> header file

int main() {
    pid_t cpid;
    int status=0;
    if (fork() == 0)
    {
        printf("Child speaking: Child exit status: %d\n", WEXITSTATUS(status));
        exit(42); // Child exits with status code 42
        }
    else
        //cpid = wait(NULL); // Parent reaps child
        cpid = wait(&status);

    //printf("Child pid: %d\n", cpid);
    //printf("Child exit status: %d\n", WEXITSTATUS(cpid));
    //printf("Child exit status: %d \n", status);
    printf("Parent speaking: Child exit status: %d\n", WEXITSTATUS(status));
    return 0;
}
