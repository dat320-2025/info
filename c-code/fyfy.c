#include <stdio.h>

int * increment_in_stack(int v){
        int c=v+1;   
        int *d=&c;
        return d;
    }
    
int main()
{
    int *a;
    printf("The pointer a is pointing to address %p and the value there is %d \n", a, *a);
    printf("increment_in_stack(10):\n");
    a=increment_in_stack(10);
    printf("The pointer a is pointing to address %p and the value there is %d \n", a, *a);
    printf("Kind of works, but not really \n");
    printf("---------------- \n");

    int *b;
    printf("The pointer b is pointing to address %p and the value there is %d \n", b, *b);
    printf("After increment_in_stack(50):\n");
    b=increment_in_stack(50);
    printf("Now the address of b is pointing to address %p and the value there is %d \n", b, *b);
    printf("Now the address of a is pointing to address %p and the value there is %d \n", a, *a);

    printf("Now the address of b is pointing to address %p and the value there is %d \n", b, *b);
    printf("WHAT !!!!!!!!!!!!\n");
    int i=0;
    while (i<10)
    {
       i++;
       printf("..\n");
    }
    printf("Now the address of a is pointing to address %p and the value there is %d \n", a, *a);
     printf("Now the address of b is pointing to address %p and the value there is %d \n", b, *b);

    printf("NO REALLY WHAT, THIS IS AN UNDEFINED BEHAVIOR !!!!!!!!!!!!\n");
    
    return 0;
}