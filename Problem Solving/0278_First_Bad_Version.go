/** 
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad 
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

 func firstBadVersion(n int) int {
    l := 0
    r := n
    h := 0
    
    for l <= r {
        m := (l+r) / 2
        
        if isBadVersion(m) == true {
            h = m
            r = m-1 
        } else {
            l = m+1
        }
    }
    return h
}