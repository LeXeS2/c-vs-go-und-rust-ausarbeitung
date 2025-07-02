#include <stdio.h>
#include <pthread.h>

#define THREAD_COUNT 5

int shared_counter = 0;
pthread_mutex_t counter_mutex;

void* increment_counter(void* arg) {
    for (int i = 0; i < 100000; ++i) {
        pthread_mutex_lock(&counter_mutex);
        arg++;
        pthread_mutex_unlock(&counter_mutex);
    }
    return arg;
}

int main() {
    pthread_t threads[THREAD_COUNT];
    pthread_mutex_init(&counter_mutex, NULL);

    for (int i = 0; i < THREAD_COUNT; ++i) {
        pthread_create(
            &threads[i], 
            NULL, 
            increment_counter, 
            (void*)&shared_counter);
    }

    for (int i = 0; i < THREAD_COUNT; ++i) {
        pthread_join(threads[i], (void**)&shared_counter);
    }

    pthread_mutex_destroy(&counter_mutex);

    printf("Finaler Wert des Counters: %d\n", shared_counter);
    return 0;
}