
const customStyles = {
    control: (provided, state) => ({
        ...provided,
        boxShadow: state.isFocused ? 'orange' : 'none !important',
    })
}
export default customStyles
