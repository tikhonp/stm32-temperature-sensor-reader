#include <stdbool.h>
#include <stdint.h>

#define MAX_SIZE 100

typedef struct {
    const uint8_t *items[MAX_SIZE];
    int front, rear;
} Queue;

void initStrQueue(Queue *q);

void enqueue(Queue *q, const uint8_t *message);

const uint8_t *dequeue(Queue *q);

const uint8_t *peek(Queue *q);

bool isEmpty(Queue *q);

bool isFull(Queue *q);

