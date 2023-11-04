#include<iostream>

// bubble sorting using cpp
void bubble_sort(int arr[], int n){
	bool sorted = true;
	while (sorted){
		sorted = false;
		for (int i = 0; i < n; i++ ) {
			if (arr[i] > arr[i+1]){
				sorted = true;
				int temp = arr[i];
				arr[i] = arr[i+1];
				arr[i+1] = temp;
			}
	}
}
}


int main() 
{
	int arr[1000];
	for (int i = 1000; i>=0; i--){
		arr[1000-i] =i ;
	}

    bubble_sort(arr,1000);

    for (int i = 0; i<1000; i++){
        std::cout << arr[i] << std::endl;
    }
	return 0;
}
