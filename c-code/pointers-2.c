#include <stdio.h>

int main()
{
int *p1;

int i=0;

p1=&i;

*p1=127;

printf("The value of i is now %d just like the value of *p1 which is %d\n",i, *p1);
printf("The address of i is now %p just like the address of p1 which is %p\n", &i, p1);

printf("Is this address the virtual or the physical address ?");


return 0;
}

