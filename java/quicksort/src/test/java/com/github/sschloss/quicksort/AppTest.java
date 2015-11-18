/*
 * QuickSort
 * Sorts in-place
 */
package com.github.sschloss.quicksort;
import junit.framework.Test;
import junit.framework.TestCase;
import junit.framework.TestSuite;
import java.util.Arrays;
import java.util.List;
import java.util.ArrayList;
import java.util.Collections;
/**
 * Unit test for simple App.
 */
public class AppTest 
    extends TestCase
{
    /**
     * Create the test case
     *
     * @param testName name of the test case
     */
    public AppTest( String testName )
    {
        super( testName );
    }

    /**
     * @return the suite of tests being tested
     */
    public static Test suite()
    {
        return new TestSuite( AppTest.class );
    }

    /**
     * Rigourous Test :-)
     */
    public void testQuickSort()
    {
		Integer[] input = { 5, 8, 4, 1, 3, 2, 7, 6 };
		Integer[] expValues = { 1, 2, 3, 4, 5, 6, 7, 8 };
		ArrayList<Integer> expected = new ArrayList();
		Collections.addAll(expected, expValues);
		QuickSort q = new QuickSort(input);
		q.sort();
		List<Integer> actual = q.getData();
		assertEquals(expected, actual);
    }
}
