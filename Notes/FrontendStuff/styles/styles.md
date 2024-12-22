# Short About CSS

## Positioning
### Static
1. The default value. 
2. Element will flow into the page as it normally would (Block, inline etc.)

### Relative
1. Positioned relative to its original position (Relative to itself)
2. Stays in its default place unless offset with `top`, `left`, `right`, or `bottom`.
3. `top: 10px` would shift the element 10px down from where it normally would be.


### Absolute
1. Position itself against the first ancestor that is **other than static**
2. `top: 0; left: 0;` would place the element in the top left corner of its first not static parent, 
or viewport if none exists.

### Fixed
1. Relative to the viewport or browser window itself.
2. Scrolling won't change its position.

### Sticky
1. Behaves like a static element until scrolled past it.
2. It then becomes fixed as long as its parent element is on the screen.
3. Check it out: https://css-tricks.com/position-sticky-2/