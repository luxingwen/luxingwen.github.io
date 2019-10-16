package com.yee;

import java.util.*;

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
        maxsubNodes(graph, 5);
    }


    // 单次扩散
    public static Map<Integer, List<Vertex>> singleDiffusion(Graph graph, List<Vertex> vertexList) {
        Random ran = new Random();
        Set<Vertex> active = new HashSet(); //存储被激活的节点
        Stack<Vertex> target = new Stack<Vertex>(); // 存储目标被激活节点
        Map<Integer, List<Vertex>> res = new HashMap<Integer, List<Vertex>>();//
        //
        for(int i = 0 ;i <vertexList.size(); i++){
            Vertex vertex = vertexList.get(i);
            target.push(vertex);
            while(target.size()>0){
                Vertex vertex1 = target.pop();
                active.add(vertex1);// 添加激活节点
                // 与它相连的节点
                List<Edge> listEdge  = graph.getG().get(vertex1);

                // 遍历与它相连的节点
                for(int j = 0; j < listEdge.size(); j++){
                    double randnum = ran.nextDouble();
                    Edge e = listEdge.get(j);
                    // 如果 有概率被激活
                    if(randnum <= e.getWeight()){
                        if(!active.contains(e.getVertex())){
                            target.push(e.getVertex());
                        }
                    }
                }

            }
            List<Vertex> list = new ArrayList<>();
            active.forEach(item -> list.add(item));
            res.put(new Integer(res.size()+1), list);
        }
        return res;
    }


    // 最大子节点
    public static void maxsubNodes(Graph graph, int k){

        List<Vertex> seedNodes = new ArrayList<Vertex>(); // 存储最大子节点集合
        Map<Vertex, Boolean>  signVertex = new HashMap<Vertex, Boolean>(); // 标记已经加载进来的节点
        seedNodes.add(new Vertex("3"));
        seedNodes.add(new Vertex("4"));
        seedNodes.add(new Vertex("8"));
        //Map<Int, Int> result = new HashMap<Int, Int>(); //will store the results

        for(int i = 0; i < k; i++){
           Map<Integer, Vertex> maxMap = new HashMap<Integer, Vertex>();
            Integer maxIndex = new Integer(0);
            Map<Vertex, List<Edge>> graphs = graph.getG();
            for(Vertex vertex : graphs.keySet()){
                if(signVertex.containsKey(vertex)){
                    continue;
                }
                List<Vertex> list = new ArrayList<Vertex>();
                list.addAll(seedNodes);
                list.add(vertex);
                Map<Integer, List<Vertex>> layers = singleDiffusion(graph, list);
                List<Vertex> listMax = layers.get(layers.size());// 获取被集火节点数最多的列表
                maxMap.put(new Integer(listMax.size()), vertex);
                maxIndex = new Integer(listMax.size());
            }
            Vertex vertexItem = maxMap.get(maxIndex);
            if(!seedNodes.contains(vertexItem)) {
                seedNodes.add(vertexItem);
            }
            signVertex.put(vertexItem, new Boolean(true));
        }

        for(int i=0; i<seedNodes.size(); i++){
            Vertex v = seedNodes.get(i);
            System.out.println(v.label);
        }
    }
}
