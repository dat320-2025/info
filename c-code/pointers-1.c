#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
int main()
{
int *a; 
int b=3;
int *lost_value=a;
a=&b;
b++;
printf("!!!!! lost pointer !!!!!!!!\n");
printf("address %p, content %p, value %d", &lost_value, lost_value, *lost_value);
return 0;
}