/**
 * Const that will define the design of the select components.
 *
 * @type {{control: (function(*, *): *&{boxShadow: string})}}
 */
const customStyles = {
    control: (provided, state) => ({
        ...provided,
        boxShadow: state.isFocused ? 'orange' : 'none !important',
    })
}
export default customStyles
