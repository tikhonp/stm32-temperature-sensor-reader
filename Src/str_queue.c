#include "str_queue.h"

void initStrQueue(Queue *q) {
    q->front = -1;
    q->rear = -1;
}

bool isFull(Queue *q) {
    return ((q->front == (q->rear + 1) % MAX_SIZE) ||
            (q->front == 0 && q->rear == MAX_SIZE - 1));
}

bool isEmpty(Queue *q) { return (q->front == -1); }

void enqueue(Queue *q, const uint8_t *message) {
    if (isFull(q)) {
        return;
    }
    if (q->front == -1)
        q->front = 0;
    q->rear = (q->rear + 1) % MAX_SIZE;
    q->items[q->rear] = message;
}

const uint8_t *dequeue(Queue *q) {
    if (isEmpty(q)) {
        static const uint8_t empty = 0;
        return &empty;
    }

    const uint8_t *message = q->items[q->front];
    if (q->front == q->rear) {
        q->front = -1;
        q->rear = -1;
    } else {
        q->front = (q->front + 1) % MAX_SIZE;
    }
    return message;
}

