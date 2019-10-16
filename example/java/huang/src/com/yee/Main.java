package com.yee;

public class Main {

    public static void main(String[] args) {

        Graph graph = new Graph();
        graph.addVertex("1");
        graph.addVertex("2");
        graph.addVertex("3");
        graph.addVertex("4");
        graph.addVertex("5");
        graph.addVertex("6");
        graph.addVertex("7");
        graph.addVertex("8");
        graph.addVertex("9");

        graph.addEdge("1", "2", 0.15);
        graph.addEdge("2", "3", 0.1);
        graph.addEdge("2", "5", 0);
        graph.addEdge("2", "4", 0.2);
        graph.addEdge("4", "5", 0.35);
        graph.addEdge("4", "7", 0.2);
        graph.addEdge("4", "9", 0.3);
        graph.addEdge("9", "7", 0.1);
        graph.addEdge("9", "8", 0.4);
        graph.addEdge("7", "5", 0.1);
        graph.addEdge("5", "6", 0.3);
        graph.printGraph();
    }
}
