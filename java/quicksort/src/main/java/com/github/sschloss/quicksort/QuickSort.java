/*
 * QuickSort
 * Sorts in-place
 */
package com.github.sschloss.quicksort;

import java.util.ArrayList;
import java.util.List;
import java.util.Collections;
public class QuickSort {

	List<Integer> data;

	public QuickSort(Integer[] input) {	
		data = new ArrayList();
		Collections.addAll(data, input);
	}

	public void sort() {
		qsort(0, data.size() - 1);
	}

	public List<Integer> getData() {
		return data; 
	}	

	private void qsort(int start, int end) {
		int n = data.size()-1;
		if ((start > n) || (end > n) || (start >= end)) {
			return;
		} 

		int i = partition(start, end);
		qsort(start, i);
		qsort(i+1, end);	
		return;
	}

	private int partition(int start, int end) {
		int pivot = data.get(start);
		int i = start + 1, j = start + 1; 
		while (j <= end) {
			if (pivot > data.get(j)) {
				swap(i, j);
				i++;
			}
			j++;
		}
		int tmp = data.get(start);
		data.set(start, data.get(i-1));
		data.set(i-1, tmp);
		return i-1;
	}	

	private void swap(int i, int j) {
		if (i < data.size() && j < data.size()) {
			int tmp = data.get(i);	
			data.set(i, data.get(j));
			data.set(j, tmp);
		}
	}
}
